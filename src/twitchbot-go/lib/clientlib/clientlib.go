package clientlib

import (
    "log"
    "os"
    "github.com/lrstanley/girc"
    "twitchbot-go/lib/configlib"
)

func New(config configlib.Config) (*girc.Client, error) {
    client := girc.New(girc.Config{
        Server: config.Irc.ServerHost,
        Port:   config.Irc.ServerPort,
        Nick:   config.Irc.ClientNickname,
        User:   config.Irc.ClientNickname,
        Debug:  os.Stdout,
    })

    // Add handlers
    client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
        log.Printf("Connected to '%s:%d'", config.Irc.ServerHost, config.Irc.ServerPort)
    })

    // Actually try to connect
    err := client.Connect()
    return client, err
}
