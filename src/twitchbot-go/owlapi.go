package main

// Program must be executed in a directory with a "teams.json" file.
// JSON schema is based on api.overwatchleague.com/teams , as of 21 September 2018.

// Program outputs information about Overwatch League teams/players to stdout,
// with some basic markdown formatting for readability.

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

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

func main() {
    var data Data

    raw_data, err := ioutil.ReadFile("teams.json")
    if err != nil {
        log.Fatalln(err)
    }

    err = json.Unmarshal(raw_data, &data)
    if err != nil {
        log.Fatalln(err)
    }

    competitors := make(map[string]Competitor)
    players     := make(map[string]Player)

    for _, competitor := range data.Competitors {
        competitors[competitor.Competitor.Name] = competitor
        fmt.Printf("**%s:**\n`", competitor.Competitor.Name)
        for _, player := range competitor.Competitor.Players {
            players[player.Player.Name] = player
            fmt.Printf("%s ", player.Player.Name)
        }
        fmt.Printf("`\n")
    }
}
