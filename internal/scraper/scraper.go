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

// ScrapeTweets scrapes tweets from Coindesk Twitter channel periodically
func ScrapeTweets(repo *db.TweetRepository) {
	ctx := context.Background()
	channel := "coindesk"
	count := 50

	scraper := twitterscraper.New().SetSearchMode(twitterscraper.SearchUsers)
	username := "JAnimate123"
	password := "kimzeey23"

	err := scraper.Login(username, password)
	if err != nil {
		panic(err)
	}

	for tweet := range scraper.GetTweets(ctx, channel, count) {
		if tweet.Error != nil {
			log.Println("Error scraping tweet:", tweet.Error)
			continue
		}

		if tweet.Tweet.ID == lastTweetID {
			log.Println("Already processed tweet:", tweet.Tweet.ID)
			continue
		}



		//get the first index in each


		// Convert tweet to our internal model
		internalTweet := model.Tweet{
			ID:        tweet.Tweet.ID,
			Text:      tweet.Tweet.Text,
		
		}
		//since we have the id lets get each tweet 



		// Convert tweet to our internal model
		if len(tweet.Photos) > 0{
			internalTweet.ImageURL = tweet.Photos[0].URL
		}
		if len(tweet.Videos) > 0 {
			internalTweet.VideoURL = tweet.Videos[0].URL
		}

		// Save the tweet to the database
		saveErr := repo.SaveTweet(internalTweet)
		if saveErr != nil {
			log.Println("Error saving tweet:", saveErr)
			continue
		}

		//check for video
		if internalTweet.VideoURL != "" {
			log.Println("Video found in tweet:", internalTweet.ID)
			email.SendEmailForVideo(internalTweet)
		}

		lastTweetID = tweet.Tweet.ID
	}
}



