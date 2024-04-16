// cmd/main.go

package main

import (
    "fmt"
	"context"

    twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
    scraper := twitterscraper.New().SetSearchMode(twitterscraper.SearchUsers)
	username := "JAnimate123"
	password := "kimzeey23"

    err := scraper.Login(username, password)
    if err != nil {
        panic(err)
    }
	for tweet := range scraper.GetTweets(context.Background(), "coindesk", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		fmt.Println(tweet.Text)
	}
}
