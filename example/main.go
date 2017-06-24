package main

import (
	"log"
	"net/url"
	"os"

	backlog "github.com/griffin-stewie/go-backlog"
)

func main() {
	token := os.Getenv("BACKLOG_TOKEN")
	if token == "" {
		log.Fatalln("You need Backlog access token.")
	}

	URL, err := url.Parse("https://YOURSPACE.backlog.jp")
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	client := backlog.NewClient(URL, token)

	issues, err := client.Issues()
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	log.Printf("issues: %#+v", issues)
}
