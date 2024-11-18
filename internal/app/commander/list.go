package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	products := c.productService.List()
	outTxMessage := "All products: \n\n"

	for _, pr := range products {
		outTxMessage += pr.Title
		outTxMessage += "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outTxMessage)
	c.bot.Send(msg)
}
