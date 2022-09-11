package main

import (
	"log"
	"os"

	"github.com/kaankivilcim/pushover"
)

func getToken() string {
	token := os.Getenv("PUSHOVER_TOKEN")
	if len(token) == 0 {
		log.Fatalln("Environment variable PUSHOVER_TOKEN was not set")
	}
	return token
}

func getMsgParams() (string, string, string) {
	if len(os.Args[1:]) != 3 {
		log.Fatalln("Incorrect number of arguments; expected user, title, msg")
	}
	return os.Args[1], os.Args[2], os.Args[3]
}

func main() {
	token := getToken()
	user, title, msg := getMsgParams()
	m := pushover.New(token, user)
	resp, err := m.Send(title, msg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Request ID: " + resp.RequestID)
}
