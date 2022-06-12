package main

import (
	"github.com/Vanek623/BabyTracker/internal/app/commands"
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
	cmdr := commands.NewCommander(bot, productService)

	for update := range updates {
		cmdr.HandleUpdate(update)
	}
}
