package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - help\n"+
			"/list- list entities",
	)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["help"] = (*Commander).Help
}
