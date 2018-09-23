package clientlib

import (
    "fmt"
    "log"
    "strings"
    "github.com/lrstanley/girc"
    "twitchbot-go/lib/owlapi"
)

func HandlePing(c *girc.Client, e girc.Event) {
    c.Cmd.Reply(e, "pong!")
}

func HandleOw(c *girc.Client, e girc.Event) {
    args := strings.Split(e.Trailing, " ")
    // args[0] == '!ow'

    // Try to find a matching team and print their members
    competitor_data_map, err := owlapi.GetCompetitorDataMap()
    if err != nil {
        log.Print(err)
        return
    }

    key := strings.ToLower(args[1])
    competitor, present := competitor_data_map[key]
    if !present {
        return
    }

    team_name := competitor.Competitor.Name
    player_names := make([]string, len(competitor.Competitor.Players))
    for i, player := range competitor.Competitor.Players {
        player_names[i] = player.Player.Name
    }
    c.Cmd.Reply(e, fmt.Sprintf("%s are: %s.", team_name, strings.Join(player_names, ", ")))

}
