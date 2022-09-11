package pushover

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	PushoverMessageAPI = "https://api.pushover.net/1/messages.json"
)

type Pushover struct {
	token string
	user  string
}

type Response struct {
	Errors    []string `json:"errors"`
	Status    int      `json:"status"`
	RequestID string   `json:"request"`
}

func New(token, user string) *Pushover {
	return &Pushover{token, user}
}

func (p *Pushover) Send(title, message string) (*Response, error) {
	r := &Response{}

	body := url.Values{
		"token":   {p.token},
		"user":    {p.user},
		"title":   {title},
		"message": {message},
	}

	resp, err := http.PostForm(PushoverMessageAPI, body)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return r, err
	}

	if len(r.Errors) > 0 {
		return r, errors.New("Pushover Message API: " + r.Errors[0])
	}

	return r, nil
}
