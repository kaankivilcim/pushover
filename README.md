# Pushover Message API Go Package

A Golang package for the [Pushover Message API](https://pushover.net/api).

## Usage

```go
package main

import (
    "log"

    "github.com/kaankivilcim/pushover"
)

func main() {
    // Create a new message for a specific user
    m := Pushover.New("api_token", "user_key")

    // Send the message
    resp, err := m.Send("title", "message")
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(resp.RequestID)
}
```