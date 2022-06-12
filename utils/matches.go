package utils

import (
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DoMatch(resp *tgbotapi.PhotoConfig, img string) {
	match, _ := regexp.MatchString(`^https://www.salonlfc.com`, img)
	if match {
		resp.Caption = "<em>Image was not found by @ABCMediaBot</em>"
	} else {
		resp.Caption = "<em>Image was found by @ABCMediaBot</em>"
	}
	resp.ParseMode = "html"
}
