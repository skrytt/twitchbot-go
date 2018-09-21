package clientlib

import (
    "fmt"
    "log"
    "strings"
    "github.com/lrstanley/girc"
    "twitchbot-go/lib/configlib"
)

func New(config configlib.Config) (*girc.Client, error) {
    supported_caps := make(map[string][]string)
    for _, capability := range(config.Irc.ClientCapabilities) {
        supported_caps[capability] = nil
    }

    client := girc.New(girc.Config{
        Server:         config.Irc.ServerHost,
        Port:           config.Irc.ServerPort,
        SSL:            true,
        ServerPass:     fmt.Sprintf("oauth:%s", config.Authorization.Token),
        Nick:           strings.ToLower(config.Irc.ClientNickname),
        User:           strings.ToLower(config.Irc.ClientNickname),
        SupportedCaps:  supported_caps,
    })

    // Add handlers
    client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
        channel_to_join := fmt.Sprintf("#%s", strings.ToLower(config.Irc.ClientChannel))
        log.Printf("Connected to '%s:%d', joining %s...",
                   config.Irc.ServerHost, config.Irc.ServerPort, channel_to_join)
        c.Cmd.Join(channel_to_join)
    })

    client.Handlers.Add(girc.JOIN, func(c *girc.Client, e girc.Event) {
        log.Printf("Joined channel '%s'", e.Params[0])

    })

    client.Handlers.Add(girc.PRIVMSG, func(c *girc.Client, e girc.Event) {
        if strings.HasPrefix(e.Trailing, "!ping") {
            c.Cmd.Reply(e, "pong!")
        }
    })

    // Actually try to connect
    err := client.Connect()
    return client, err
}
