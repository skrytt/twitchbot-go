package main

// Program outputs information about Overwatch League teams/players to stdout,
// sourced from Overwatch League API.
// Some basic markdown formatting is used for readability.

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
