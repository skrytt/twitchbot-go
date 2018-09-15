package clientlib

import (
    "fmt"
    "log"
    "os"
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
        ServerPass:     fmt.Sprintf("oauth:%s", config.Irc.Token),
        Nick:           strings.ToLower(config.Irc.ClientNickname),
        User:           strings.ToLower(config.Irc.ClientNickname),
        Debug:          os.Stdout,
        SupportedCaps:  supported_caps,
    })

    // Add handlers
    client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
        log.Printf("Connected to '%s:%d'", config.Irc.ServerHost, config.Irc.ServerPort)

        // Try to join the configured channel
        // TODO
    })

    // Actually try to connect
    err := client.Connect()
    return client, err
}
