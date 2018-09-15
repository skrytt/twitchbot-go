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
    }
    Irc                 struct {
        ServerHost          string
        ServerPort          int
        Token               string
        ClientCapabilities  []string
        ClientNickname      string
        ClientChannel       string
    }
}

// loadConfig: Load a JSON config from disk.
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

