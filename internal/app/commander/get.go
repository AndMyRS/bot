package commander

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {

	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := c.productService.Get(idx)

	if err != nil {
		log.Printf("failed to get product with id [%d]: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, product.Title)
	c.bot.Send(msg)
}
