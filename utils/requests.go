package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/Muhammadali-Akbarov/telebot-golang/models"
)

func SendRequest(search string) string {
	var mystruct models.SimpleStruct
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
