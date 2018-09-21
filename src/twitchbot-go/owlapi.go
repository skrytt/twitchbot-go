package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

// Teams
type Data struct {
    Id                  int
    AvailableLanguages  []string
    Competitors         []struct {
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
            Players             []struct {
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
                Flags               []string
                Team                struct {
                    Id                  int
                    Type                string
                }
            }
            PrimaryColor        string
            SecondaryColor      string
            SecondaryPhoto      string
            Type                string
        }
    }
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

type Player struct {
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

    fmt.Printf("%+v", data)

}
