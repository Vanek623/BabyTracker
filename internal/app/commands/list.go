package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(messageIn *tgbotapi.Message) {
	text := ""

	for _, p := range c.productService.List() {
		text += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(messageIn.Chat.ID, text)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "list_10"),
		),
	)

	c.bot.Send(msg)
}
