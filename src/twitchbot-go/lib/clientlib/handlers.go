package clientlib

import (
    "github.com/lrstanley/girc"
)

func HandlePing(c *girc.Client, e girc.Event) {
    c.Cmd.Reply(e, "pong!")
}
