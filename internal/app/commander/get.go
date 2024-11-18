package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "TBD")
	c.bot.Send(msg)
}
