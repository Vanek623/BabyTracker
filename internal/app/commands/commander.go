package commands

import (
	"fmt"
	"github.com/Vanek623/BabyTracker/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
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

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	if update.CallbackQuery != nil {
		args := strings.Split(update.CallbackQuery.Data, "_")
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Command: %s\n", args[0])+
				fmt.Sprintf("Offset: %s", args[1]))

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
