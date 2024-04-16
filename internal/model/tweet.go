package model

import "time"

type Tweet struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	ImageURL  string    `json:"image_url"`
	VideoURL  string    `json:"video_url"`
}
