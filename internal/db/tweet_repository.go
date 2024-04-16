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

// TweetExists checks if a tweet with the given ID exists in the database
func (repo *TweetRepository) TweetExists(tweetID string) bool {
	var exists bool
	err := repo.db.QueryRow("SELECT EXISTS(SELECT 1 FROM tweets WHERE tweet_id = $1)", tweetID).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if tweet exists: %v", err)
		return false
	}
	return exists
}


// GetAllTweets retrieves all tweets from the database
func (repo *TweetRepository) GetAllTweets() ([]model.Tweet, error) {
	var tweets []model.Tweet

	rows, err := repo.db.Query("SELECT tweet_id, text, created_at, image_url, video_url FROM tweets")
	if err != nil {
		return nil, fmt.Errorf("error querying tweets: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tweet model.Tweet
		err := rows.Scan(&tweet.ID, &tweet.Text, &tweet.CreatedAt, &tweet.ImageURL, &tweet.VideoURL)
		if err != nil {
			return nil, fmt.Errorf("error scanning tweet row: %v", err)
		}
		tweets = append(tweets, tweet)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tweets: %v", err)
	}

	return tweets, nil
}

// GetTweetsWithPagination retrieves tweets with pagination from the database
func (repo *TweetRepository) GetTweetsWithPagination(page, limit int) ([]model.Tweet, error) {
	var tweets []model.Tweet

	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT tweet_id, text, created_at, image_url, video_url FROM tweets LIMIT %d OFFSET %d", limit, offset)
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying tweets with pagination: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tweet model.Tweet
		err := rows.Scan(&tweet.ID, &tweet.Text, &tweet.CreatedAt, &tweet.ImageURL, &tweet.VideoURL)
		if err != nil {
			return nil, fmt.Errorf("error scanning tweet row: %v", err)
		}
		tweets = append(tweets, tweet)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tweets: %v", err)
	}

	return tweets, nil
}
