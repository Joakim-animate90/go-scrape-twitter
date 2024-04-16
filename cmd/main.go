package main

import (
	"database/sql"
	"log"
	//"time"

	"github.com/Joakim-animate90/go-scrape-twitter/internal/db"
	"github.com/Joakim-animate90/go-scrape-twitter/internal/scraper"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize database connection
	dbConn, err := sql.Open("postgres", "postgres://postgres:kimzeey23@localhost:5432/twitter_scraper?sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	// Create a TweetRepository instance
	repo := db.NewTweetRepository(dbConn)

	// Run the scraper periodically
	//ticker := time.NewTicker(1 * time.Hour) // Adjust the interval as needed
	//defer ticker.Stop()

	//for range ticker.C {
		scraper.ScrapeTweets(repo)
	//}
}
