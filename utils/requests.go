package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/Muhammadali-Akbarov/telebot-golang/config"
	"github.com/Muhammadali-Akbarov/telebot-golang/models"
)

func SendRequest(search string) string {
	var mystruct models.SimpleStruct
	api_key := config.GetEnv("API_KEY")
	api_host := config.GetEnv("API_HOST")
	img_404 := config.GetEnv("IMG_404")

	url := fmt.Sprintf("https://%v/images/search?q=%v", api_host, search)
	client := &http.Client{}
	r, errReq := http.NewRequest("GET", url, nil)

	if errReq != nil {
		log.Println(errReq)

	}

	r.Header.Add("X-RapidAPI-Key", api_key)
	r.Header.Add("X-RapidAPI-Host", api_host)

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
		image := img_404
		return image
	} else {
		myIndex := rand.Intn(len(mystruct.Value))
		image := fmt.Sprintf("%v", mystruct.Value[myIndex])
		replaced := regexp.MustCompile(`^{+|}+$`).ReplaceAllString(image, ``)
		return replaced
	}
}
