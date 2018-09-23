package owlapi

// Program must be executed in a directory with a "teams.json" file.
// JSON schema is based on api.overwatchleague.com/teams , as of 21 September 2018.

// Program outputs information about Overwatch League teams/players to stdout,
// with some basic markdown formatting for readability.

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
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

func GetCompetitorDataMap() (map[string]Competitor, error) {
    if _competitor_data_map == nil {
        new_competitor_data_map := make(map[string]Competitor)

        api_data, err := GetApiData()
        if err != nil {
            return new_competitor_data_map, err
        }

        for _, competitor := range api_data.Data.Competitors {
            new_competitor_data_map[competitor.Competitor.Name] = competitor
        }

        _competitor_data_map = &new_competitor_data_map
    }
    return *_competitor_data_map, nil
}
