package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	botToken := os.Getenv("BOT_API")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Set the bot API instance's debugging mode
	//    bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up a new updates channel to receive updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	// Set up a handler function to be called on each update
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// If the update is a command, handle it
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				// Send a message to the user
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "/google")
				bot.Send(msg)
			case "google":
				pictureData := GetPicture()
				fileBytes := tgbotapi.FileBytes{Name: "google.jpg", Bytes: pictureData}
				photoConfig := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, fileBytes)
				photoConfig.Caption = "Скрин гугла"
				bot.Send(photoConfig)
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(msg)
		}
	}
}

func GetPicture() []byte {
	url := "http://go_screen:80/picture"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Use a buffer to store the response body
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body
	return buf.Bytes()
}
