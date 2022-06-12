package main

import (
	"fmt"
	"log"

	"github.com/Muhammadali-Akbarov/telebot-golang/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("5025387786:AAEpYYZySyJ0SGMDn0DFhKMxpyleJFx3aBM")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if err != nil {
			log.Println(err)
		}
		img, err := utils.GetImage(fmt.Sprintf("%v", update.Message.Text))
		if err != nil {
			log.Println(err)
		}
		file := tgbotapi.FileURL(img)
		resp_image := tgbotapi.NewPhoto(int64(update.Message.Chat.ID), file)
		utils.DoMatch(&resp_image, img)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Keyboard Mode")
		greeting := fmt.Sprintf("Welcome <b>%v</b>", update.Message.Chat.FirstName)
		to_admin := tgbotapi.NewMessage(update.Message.Chat.ID, greeting)
		to_admin.ParseMode = "html"
		switch update.Message.Text {
		case "/start":
			bot.Send(to_admin)
		case "open":
			msg.ReplyMarkup = utils.NumericKeyboard
			bot.Send(msg)
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
		default:
			if _, err := bot.Send(resp_image); err != nil {
				log.Panic(err)
			}
		}
	}
}
