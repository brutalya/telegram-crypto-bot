package tgbot

import (
	"fmt"
	"log"
	"strings"

	"github.com/brutalya/telegram-crypto-bot/crypto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CommanHandler handles user commands
func CommanHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Extract the user's command
	command := update.Message.Command()

	// Check if the command is "/price"
	if command == "price" {
		// Extract the cryptocurrency symbol from the user's message
		parts := strings.Fields(update.Message.Text)
		if len(parts) < 2 {
			sendMessage(bot, update.Message.Chat.ID, "Please specify a cryptocurrency symbol.")
			return
		}
		symbol := parts[1]

		// Get the cryptocurrency price
		price, err := crypto.GetCryptoPrice(symbol)
		if err != nil {
			sendMessage(bot, update.Message.Chat.ID, "Error getting cryptocurrency price.")
			return
		}

		// Send the cryptocurrency price to the user
		message := fmt.Sprintf("%s price: $%.2f", strings.ToUpper(symbol), price)
		sendMessage(bot, update.Message.Chat.ID, message)
	}
}

// sendMessage sends a message to the specified chat.
func sendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
