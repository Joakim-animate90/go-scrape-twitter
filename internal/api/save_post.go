package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Joakim-animate90/go-scrape-twitter/internal/db"
	_ "github.com/lib/pq"
	"log"
	

)

func GetSavedPostsHandler( repo *db.TweetRepository ,w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters from query string
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
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

	response, err := json.Marshal(tweets)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(response)
}
