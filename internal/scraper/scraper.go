package scraper

import (
	"context"
	"log"

	"github.com/Joakim-animate90/go-scrape-twitter/internal/db"
	email "github.com/Joakim-animate90/go-scrape-twitter/internal/email"
	"github.com/Joakim-animate90/go-scrape-twitter/internal/model"
	twitterscraper "github.com/n0madic/twitter-scraper"

	_ "github.com/lib/pq"
)

var lastTweetID string

func ScrapeTweets(repo *db.TweetRepository) {
    ctx := context.Background()
    channel := "coindesk"
    count := 50

    scraper := twitterscraper.New().SetSearchMode(twitterscraper.SearchUsers)
    username := ""
    password := ""

    err := scraper.Login(username, password)
    if err != nil {
        panic(err)
    }

    for tweet := range scraper.GetTweets(ctx, channel, count) {
        if tweet.Error != nil {
            log.Println("Error scraping tweet:", tweet.Error)
            continue
        }

        if repo.TweetExists(tweet.Tweet.ID) {
            log.Println("Tweet already exists in the database:", tweet.Tweet.ID)
            continue
        }

        internalTweet := model.Tweet{
            ID:   tweet.Tweet.ID,
            Text: tweet.Tweet.Text,
        }

        if len(tweet.Photos) > 0 {
            internalTweet.ImageURL = tweet.Photos[0].URL
        }
        if len(tweet.Videos) > 0 {
            internalTweet.VideoURL = tweet.Videos[0].URL
        }

        saveErr := repo.SaveTweet(internalTweet)
        if saveErr != nil {
            log.Println("Error saving tweet:", saveErr)
            continue
        }

        if internalTweet.VideoURL != "" {
            log.Println("Video found in tweet:", internalTweet.ID)
            email.SendEmailForVideo(internalTweet)
        }

        lastTweetID = tweet.Tweet.ID
    }
}
