package scraper


import (
	"context"
	"log"
	"github.com/n0madic/twitter-scraper"
	"github.com/Joakim-animate90/go-scrape-twitter/model"
)

var lastTweetID string

func scrapeTweets() {
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

		// Check if this tweet is already processed
		if tweet.Tweet.ID == lastTweetID {
			log.Println("Already processed tweet:", tweet.Tweet.ID)
			continue
		}

		// Convert tweet to our internal model
		internalTweet := model.Tweet{
			ID:        tweet.Tweet.ID,
			Text:      tweet.Tweet.Text,
			CreatedAt: tweet.Tweet.CreatedAt,
			ImageURL:  tweet.Tweet.ImageURL,
			VideoURL:  tweet.Tweet.VideoURL,
		}

		// Save the tweet to the database
		saveTweetToDB(internalTweet)

		// Check for video and send email
		sendEmailForVideo(internalTweet)

		// Update lastTweetID
		lastTweetID = tweet.Tweet.ID
	}
}





// ScrapeTweetsPeriodically scrapes tweets from Coindesk Twitter channel periodically
//func ScrapeTweetsPeriodically() {
	//for {
		// Scrape tweets from Coindesk Twitter channel
		//scraper := twitterscraper.New()
		//err := scraper.LoginOpenAccount()
		//if err != nil {
			//panic(err)
		//}
		//for tweet := range scraper.GetTweets(context.Background(), "twitter", 50) {
			//if tweet.Error != nil {
			//	panic(tweet.Error)
			//}
			//fmt.Println(tweet.Text)
		//}
		// Sleep for 1 hour before scraping again
		//time.Sleep(time.Hour)
	//}
//}

