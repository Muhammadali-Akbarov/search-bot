package app

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SayWelcome(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	welcome := fmt.Sprintf("Welcome <b>%v</b>", update.Message.Chat.FirstName)
	to_client := tgbotapi.NewMessage(update.Message.Chat.ID, welcome)
	to_client.ParseMode = "html"
	return &to_client
}
