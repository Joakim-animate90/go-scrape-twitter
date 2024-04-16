package email

import (
	"log"
	"github.com/Joakim-animate90/go-scrape-twitter/internal/model"
	"gopkg.in/gomail.v2"
)

func SendEmailForVideo(tweet model.Tweet) {
	// Check if the tweet contains a video
	if tweet.VideoURL != "" {
		// Compose email message
		subject := "Video Found in Tweet"
		body := "The following tweet contains a video:\n\n" +
			"Tweet ID: " + tweet.ID + "\n" +
			"Tweet Text: " + tweet.Text + "\n" +
			"Video URL: " + tweet.VideoURL

		// Send email
		err := sendEmail(subject, body)
		if err != nil {
			log.Println("Error sending email:", err)
			return
		}
		log.Println("Email sent successfully for tweet with video:", tweet.ID)
	}
}

func sendEmail(subject, body string) error {
	// Email configuration
	sender := "joakimbwire23@gmail.com"
	password := "hawx alvt nfsb xsnv"
	recipient := "joakimngeso@gmail.com"

	// Compose email message
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Create a new SMTP client session
	d := gomail.NewDialer("smtp.gmail.com", 465, sender, password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
