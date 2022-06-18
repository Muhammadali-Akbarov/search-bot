package models

import (
	"fmt"

	"github.com/Muhammadali-Akbarov/telebot-golang/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type TelegramUser struct {
	gorm.Model
	Tid      string `json:"tid"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&TelegramUser{})
}

func (b *TelegramUser) RegisterUser(Id int64) *TelegramUser {
	var user TelegramUser
	res := db.Where("TID=?", Id).Find(&user)

	if res.Error != nil {
		db.NewRecord(b)
		fmt.Println("works")
		err := db.Create(&b)
		fmt.Println("error:", err.Error)
		return b
	}
	return nil
}
