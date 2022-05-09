package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func main() {

	bot, err := tgbotapi.NewBotAPI("5368494360:AAGHmkcn_CCzzKXYTfVUmoPtBwcqeQOkG0Y")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		if err != nil {
			log.Println(err)
		}
		search := fmt.Sprintf("%v", update.Message.Text)
		search = strings.ReplaceAll(search, " ", "")
		img := sendRequest(search)
		file := tgbotapi.FileURL(img)
		resp_image := tgbotapi.NewPhoto(int64(update.Message.Chat.ID), file)
		match, _ := regexp.MatchString(`^https://www.salonlfc.com`, img)
		if match {
			resp_image.Caption = "<em>Image was not found by @SearchImagesRobot</em>"
		} else {
			resp_image.Caption = "<em>Image was found by @SearchImagesRobot</em>"
		}
		resp_image.ParseMode = "html"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Keyboard Mode")
		greeting := fmt.Sprintf("Welcome <b>%v</b>", update.Message.Chat.FirstName)
		to_admin := tgbotapi.NewMessage(update.Message.Chat.ID, greeting)
		to_admin.ParseMode = "html"
		switch update.Message.Text {
		case "/start":
			bot.Send(to_admin)
		case "open":
			msg.ReplyMarkup = numericKeyboard
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

type SimpleStruct struct {
	Value []struct {
		ThumbnailUrl string `json:"thumbnailUrl"`
	}
}

func sendRequest(search string) string {
	var mystruct SimpleStruct
	url := fmt.Sprintf("https://bing-image-search1.p.rapidapi.com/images/search?q=%v", search)
	client := &http.Client{}
	r, errReq := http.NewRequest("GET", url, nil)

	if errReq != nil {
		log.Println(errReq)

	}

	r.Header.Add("X-RapidAPI-Host", "bing-image-search1.p.rapidapi.com")
	r.Header.Add("X-RapidAPI-Key", "ff6bdaf671mshcc8f1ee56cc16cbp1b5c71jsnde0483b0b93d")

	resp, _ := client.Do(r)
	bytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	unMarshall := json.Unmarshal(bytes, &mystruct)
	if unMarshall != nil {
		log.Println(unMarshall)
	}

	if len(mystruct.Value) == 0 {
		image := "https://www.salonlfc.com/wp-content/uploads/2018/01/image-not-found-scaled-1150x647.png"
		return image
	} else {
		myIndex := rand.Intn(len(mystruct.Value))
		image := fmt.Sprintf("%v", mystruct.Value[myIndex])
		replaced := regexp.MustCompile(`^{+|}+$`).ReplaceAllString(image, ``)
		return replaced
	}
}
