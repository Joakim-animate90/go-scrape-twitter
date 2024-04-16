package email

import (
    "log"
    "net/smtp"
    "strconv"
    "github.com/Joakim-animate90/go-scrape-twitter/internal/model"
)

func sendEmailForVideo(tweet model.Tweet) {
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
    smtpHost := "smtp.example.com"
    smtpPort := 587
    sender := "your-email@example.com"
    password := "your-email-password"
    recipient := "recipient@example.com"

    // Compose email message
    message := "From: " + sender + "\n" +
        "To: " + recipient + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    // Send email using SMTP
    auth := smtp.PlainAuth("", sender, password, smtpHost)
    err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, sender, []string{recipient}, []byte(message))
    if err != nil {
        return err
    }
    return nil
}
