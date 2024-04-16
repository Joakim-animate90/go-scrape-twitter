package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Joakim-animate90/go-scrape-twitter/internal/api"
	"github.com/Joakim-animate90/go-scrape-twitter/internal/db"
	"github.com/Joakim-animate90/go-scrape-twitter/internal/scraper"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
)

func main() {
	// Initialize database connection
	dbUsername := ""
	dbPassword := ""
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "twitter_scraper"

	// Initialize database connection
	dbConn, err := sql.Open("postgres", "postgres://"+dbUsername+":"+dbPassword+"@"+dbHost+":"+dbPort+"/"+dbName+"?sslmode=disable")

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	// Create a TweetRepository instance
	repo := db.NewTweetRepository(dbConn)

	// Create a new cron job scheduler
	c := cron.New()

	// Define the cron job to run the scraping process every hour
	_, err = c.AddFunc("@hourly", func() {
		scraper.ScrapeTweets(repo)
		log.Println("Scraping process executed at", time.Now().Format(time.RFC3339))
	})
	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}

	// Start the cron scheduler
	c.Start()

	// Define HTTP handlers
	http.HandleFunc("/api/saved-posts", func(w http.ResponseWriter, r *http.Request) {
		api.GetSavedPostsHandler(repo, w, r)
	})
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs"))))

	// Start HTTP server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
