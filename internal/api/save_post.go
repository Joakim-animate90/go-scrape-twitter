package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Joakim-animate90/go-scrape-twitter/internal/db"
	_ "github.com/lib/pq"
	"log"
	

)
// @Summary Get saved posts
// @Description Get all saved posts from the database
// @Tags saved-posts
// @Produce json
// @Success 200 {array} Post
//@Router /api/saved-posts [get]
func GetSavedPostsHandler( repo *db.TweetRepository ,w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters from query string
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Default values if parameters are not provided or invalid
	page := 1
	limit := 10

	// Parse page number
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	// Parse limit
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}
	//log page and limit
	log.Println("Page:", page, "Limit:", limit)





	tweets, err := repo.GetTweetsWithPagination(page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch tweets", http.StatusInternalServerError)
		return
	}

	// Marshal the tweets slice to JSON
	response, err := json.Marshal(tweets)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write JSON response
	w.Write(response)
}
