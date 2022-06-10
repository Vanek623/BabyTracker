package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Default(messageIn *tgbotapi.Message) {
	//log.Printf("[%s] %s", messageIn.From.UserName, messageIn.Text)

	msg := tgbotapi.NewMessage(messageIn.Chat.ID, "Я люблю Аню!\nАня - лучшая жена!")

	//msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}
