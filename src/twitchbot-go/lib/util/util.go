package util

import (
    "log"
    "net/http"
    "twitchbot-go/lib/configlib"
)

func PrintAuthorizationUrl(c configlib.Config) error {
    req, err := http.NewRequest("GET", c.Authorization.BaseUrl, nil)
    if err != nil {
        return err
    }

    // Update the query string to include the configured auth params
    q := req.URL.Query()
    q.Add("client_id", c.Authorization.ClientId)
    q.Add("redirect_uri", c.Authorization.RedirectUri)
    q.Add("response_type", "token")
    q.Add("scope", "chat_login")
    req.URL.RawQuery = q.Encode()

    log.Println(req.URL.String())
    return nil
}
