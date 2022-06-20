package commands

import (
	"encoding/json"
	"fmt"
	"github.com/Vanek623/BabyTracker/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, ps *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: ps,
	}
}

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsedData),
		)
		c.bot.Send(msg)

		return
	}

	if update.Message == nil { // If we got a message
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
