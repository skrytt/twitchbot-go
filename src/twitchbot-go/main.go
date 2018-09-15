package main

import (
    "log"
    "twitchbot-go/lib/configlib"
    "twitchbot-go/lib/clientlib"
    "twitchbot-go/lib/util"
)

func main() {
    config_file := "config/config.json"
    config, err := configlib.Load(config_file)
    if err != nil {
        log.Fatalf("Could not load config from '%s', reason: %s", config_file, err)
    }

    util.PrintAuthorizationUrl(config)

    if config.Irc.Token == "" {
        log.Fatalf("Exiting because Irc.Token needs to be set in the config.")
    }

    client, err := clientlib.New(config)
    if err != nil {
        log.Fatalf("Could not connect to IRC server")
    }

    log.Println("Success (client %r)", client)
}
