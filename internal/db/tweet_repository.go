// internal/db/tweet_repository.go
package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Joakim-animate90/go-scrape-twitter/internal/model"
	_ "github.com/lib/pq"
)

// TweetRepository handles interactions with the tweet table in the database
type TweetRepository struct {
	db *sql.DB
}

// NewTweetRepository initializes a new TweetRepository
func NewTweetRepository(db *sql.DB) *TweetRepository {
	return &TweetRepository{db: db}
}

// SaveTweet saves a tweet to the database
func (repo *TweetRepository) SaveTweet(tweet model.Tweet) error {
	// Prepare the INSERT statement
	stmt, err := repo.db.Prepare("INSERT INTO tweets (tweet_id, text, created_at, image_url, video_url) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return fmt.Errorf("error preparing SQL statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(tweet.ID, tweet.Text, tweet.CreatedAt, tweet.ImageURL, tweet.VideoURL)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %v", err)
	}

	log.Println("Tweet saved to the database successfully")
	return nil
}
