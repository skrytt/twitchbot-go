package configlib

import (
    "encoding/json"
    "io/ioutil"
)

type Config struct {
    Authorization       struct {
        BaseUrl             string
        ClientId            string
        RedirectUri         string
        Token               string
    }
    Irc                 struct {
        ServerHost          string
        ServerPort          int
        ClientCapabilities  []string
        ClientNickname      string
        ClientChannel       string
    }
}

// Load: Load a JSON config from disk.
func Load(path string) (Config, error) {
    var config Config

    raw_data, err := ioutil.ReadFile(path)
    if err != nil {
        return config, err
    }

    err = json.Unmarshal(raw_data, &config)
    if err != nil {
        return config, err
    }

    return config, nil
}

