package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(messageIn *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(messageIn.Chat.ID, "/help\n/list")

	c.bot.Send(msg)
}
