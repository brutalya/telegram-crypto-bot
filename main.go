package main

import (
	"log"
	"os"

	"github.com/brutalya/telegram-crypto-bot/tgbot"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the Telegram bot token from the environment variable
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Initialize the Telegram bot with the bot token
	bot, err := tgbot.InitializeBot(botToken)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Start polling for updates and responding to user messages
	tgbot.StartBotPolling(bot)

	// ... Other code for your project
}
