package main

// Program must be executed in a directory with a "teams.json" file.
// JSON schema is based on api.overwatchleague.com/teams , as of 21 September 2018.

// Program outputs information about Overwatch League teams/players to stdout,
// with some basic markdown formatting for readability.

import (
    "fmt"
    "log"
    "twitchbot-go/lib/owlapi"
)

func main() {
    competitors, err := owlapi.GetCompetitorDataMap()
    if err != nil {
        log.Fatalln(err)
    }

    for _, competitor := range competitors {
        fmt.Printf("**%s:**\n`", competitor.Competitor.Name)
        for _, player := range competitor.Competitor.Players {
            fmt.Printf("%s ", player.Player.Name)
        }
        fmt.Printf("`\n")
    }
}
