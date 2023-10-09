package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// InitializeBot initializes and returns a Telegram bot instance.
func InitializeBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

// StartBotPolling starts polling for updates from the Telegram bot.
func StartBotPolling(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
		return
	}

	// for update := range updates {
	// 	if update.Message == nil {
	// 		continue
	// 	}

	// 	// Read user's message
	// 	userMessage := update.Message.Text

	// 	// Process the user's message here and send responses
	// 	// For example, you can use a switch statement to handle different commands or responses

	// 	// Echo the user's message as a response
	// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You said: "+userMessage)
	// 	_, err := bot.Send(msg)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Handle user commands
		if update.Message.IsCommand() {
			CommanHandler(bot, update)
			continue
		}
	}
}
