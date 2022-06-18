package utils

import (
	"fmt"

	"github.com/Muhammadali-Akbarov/telebot-golang/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RegisterUserFunc(chat tgbotapi.Update) {
	user := models.TelegramUser{
		Tid:      fmt.Sprintf("%v", chat.Message.Chat.ID),
		Name:     fmt.Sprintf("%v", chat.Message.Chat.FirstName),
		Username: fmt.Sprintf("%v", chat.Message.Chat.UserName),
	}
	user.RegisterUser(chat.Message.Chat.ID)
}
