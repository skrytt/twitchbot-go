package owlapi

// Program must be executed in a directory with a "teams.json" file.
// JSON schema is based on api.overwatchleague.com/teams , as of 21 September 2018.

// Program outputs information about Overwatch League teams/players to stdout,
// with some basic markdown formatting for readability.

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "time"
)

const max_cached_api_data_age time.Duration = 3600 * time.Second;
var _api_data_handle *ApiDataHandle;
var _competitor_data_map *map[string]Competitor

func invalidateCachedData() {
    _api_data_handle = nil
    _competitor_data_map = nil
}

// Teams

type Player struct {
    Flags               []string
    Team                struct {
        Id                  int
        Type                string
    }
    Player              struct {
        Id                  int
        Erased              bool
        Handle              string
        Name                string
        Type                string
        FamilyName          string
        AvailableLanguages  []string
        AttributesVersion   string
        Game                string
        Accounts            []struct {
            Id              int
            CompetitorId    int
            AccountType     string
            Value           string
            IsPublic        bool
        }
        Headshot            string
        Nationality         string
        Attributes          struct {
            HeroImage           string
            Heroes              []string
            Hometown            string
            Player_Number       int
            Preferred_Slot      string
            Role                string
        }
        GivenName           string
        HomeLocation        string
    }
}

type Competitor struct {
    Division            struct {
        Id                  int
    }
    Competitor          struct {
        Id                  int
        AbbreviatedName     string
        Accounts            []struct {
            CompetitorId        int
            IsPublic            bool
            AccountType         string
            Id                  int
            Value               string
        }
        AddressCountry      string
        Attributes          struct {
            Team_Guid           string
            City                string
            Manager             string
            Hero_Image          string
        }
        AttributesVersion   string
        AvailableLanguages  []string
        Game                string
        Handle              string
        HomeLocation        string
        Icon                string
        Logo                string
        Name                string
        OwlDivision         int
        Players             []Player
        PrimaryColor        string
        SecondaryColor      string
        SecondaryPhoto      string
        Type                string
    }
}

type Aliases struct {
    Teams   map[string][]string
}

type Data struct {
    Id                  int
    AvailableLanguages  []string
    Competitors         []Competitor
    Description         string
    Game                string
    Logo                string
    Name                string
    OwlDivisions        []struct {
        Id                  int
        Name                string
        String              string
    }
    Strings             interface{}
}

type ApiDataHandle struct {
    Data        Data
    Timestamp   time.Time
}

func GetApiData() (*ApiDataHandle, error) {
    // We cache the API data to avoid repeatedly downloading it.
    // Check if our cached version should be updated
    if _api_data_handle == nil || time.Since(_api_data_handle.Timestamp) > max_cached_api_data_age {
        // We'll download a new version of the data
        client := &http.Client{
            Timeout: 10 * time.Second,
        }

        response, err := client.Get("https://api.overwatchleague.com/teams")
        if err != nil {
            return nil, err
        }

        bytes_data, err := ioutil.ReadAll(response.Body)
        if err != nil {
            return nil, err
        }

        new_api_data_handle := ApiDataHandle{}
        new_api_data := &(new_api_data_handle.Data)

        err = json.Unmarshal(bytes_data, &new_api_data)
        if err != nil {
            return nil, err
        }
        new_api_data_handle.Timestamp = time.Now()

        // Invalidate caches to force regeneration of data maps
        invalidateCachedData()
        _api_data_handle = &new_api_data_handle
    }
    return _api_data_handle, nil
}

// Note: the OWL API refers to teams as "competitors",
// so we use the same convention here.
func GetCompetitorDataMap() (map[string]Competitor, error) {

    // Initialize the mapping if required.
    if _competitor_data_map == nil {
        new_competitor_data_map := make(map[string]Competitor)

        // Since competitors are called by nicknames that aren't always present
        // in the API, we use a configuration file to determine the
        // mapping of nicknames to competitor objects.
        var aliases Aliases
        config_file := "config/aliases.json"
        raw_alias_data, err := ioutil.ReadFile(config_file)
        if err != nil {
            return new_competitor_data_map, err
        }
        err = json.Unmarshal(raw_alias_data, &aliases)
        if err != nil {
            return new_competitor_data_map, err
        }

        api_data, err := GetApiData()
        if err != nil {
            return new_competitor_data_map, err
        }

        log.Printf("Building map of competitors...")
        for _, competitor := range api_data.Data.Competitors {

            // Use the lowercased competitor names as a key to look up aliases..
            competitor_name := competitor.Competitor.Name
            competitor_name_lc := strings.ToLower(competitor_name)

            competitor_aliases, present := aliases.Teams[competitor_name_lc]
            if present {
                log.Printf("Mapping aliases for '%s'", competitor_name)
                for _, competitor_alias := range competitor_aliases {
                    competitor_alias_lc := strings.ToLower(competitor_alias)
                    new_competitor_data_map[competitor_alias_lc] = competitor
                }
            } else {
                log.Printf("No aliases configured for '%s'", competitor_name)
            }
        }

        _competitor_data_map = &new_competitor_data_map
    }
    return *_competitor_data_map, nil
}
