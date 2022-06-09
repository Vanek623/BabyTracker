package main

import (
	"github.com/Vanek623/BabyTracker/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil { // If we got a message
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBeh(bot, update.Message)

			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, messageIn *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(messageIn.Chat.ID, "/help\n/list")

	bot.Send(msg)
}

func defaultBeh(bot *tgbotapi.BotAPI, messageIn *tgbotapi.Message) {
	//log.Printf("[%s] %s", messageIn.From.UserName, messageIn.Text)

	msg := tgbotapi.NewMessage(messageIn.Chat.ID, "Я люблю Аню!\nАня - лучшая жена!")

	//msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, messageIn *tgbotapi.Message, productService *product.Service) {
	text := ""

	for _, p := range productService.List() {
		text += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(messageIn.Chat.ID, text)

	bot.Send(msg)
}
