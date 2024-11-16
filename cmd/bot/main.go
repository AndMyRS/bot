package main

import (
	"log"
	"os"

	"github.com/AndMyRs/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil { // If we got a message

			switch update.Message.Command() {
			case "help":
				HelpCommand(bot, update.Message)
			case "list":
				ListCommand(bot, update.Message, productService)
			default:
				defaultActions(bot, update.Message)
			}

		}
	}
}

func HelpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - help\n"+
			"/list- list entities",
	)
	bot.Send(msg)
}

func defaultActions(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}

func ListCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message, productService *product.Service) {
	products := productService.List()
	outTxMessage := "All products: \n\n"

	for _, pr := range products {
		outTxMessage += pr.Title
		outTxMessage += "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outTxMessage)
	bot.Send(msg)
}
