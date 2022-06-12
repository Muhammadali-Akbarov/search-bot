package main

import (
	"fmt"
	"log"

	"github.com/Muhammadali-Akbarov/telebot-golang/app"
	"github.com/Muhammadali-Akbarov/telebot-golang/config"
	"github.com/Muhammadali-Akbarov/telebot-golang/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	updates, bot, err := app.Start(config.GetEnv("API_TOKEN"))
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

		switch update.Message.Text {
		case "/start":
			bot.Send(app.SayWelcome(&update))
		default:
			if _, err := bot.Send(resp_image); err != nil {
				log.Panic(err)
			}
		}
	}
}
