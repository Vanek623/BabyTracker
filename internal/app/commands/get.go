package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()

	arg, err := strconv.Atoi(args)
	var str string
	if err != nil {
		str = fmt.Sprintf("wrong args %v", args)
	} else {
		p, err := c.productService.Get(arg)
		if err != nil {
			str = err.Error()
		} else {
			str = fmt.Sprintf(p.Title)
		}
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, str)
	c.bot.Send(msg)
}
