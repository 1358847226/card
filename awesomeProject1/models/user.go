package models

import "time"

type User struct {
	Id       int64  `gorm:"column:user_id" json:"id"`
	Account	 string	`gorm:"column:account" json:"account"`
	// userId
	Name     string `gorm:"column:user_name" json:"name"`
	User_password string `gorm:"column:user_password" json:"user_password"`
	Power    int64   `gorm:"column:user_power" json:"power"`
	Share    int64  `gorm:"column:user_share" json:"share"`
	Createdate time.Time `gorm:"column:createdate" json:"createdate"`
	Admin_id	string `json:"admin_id"`
	Avatar string `json:"avatar"`
	Card_quantity int `json:"card_quantity"`
	Max_quantity int `json:"max_quantity"`
	Roleid     string `json:"roleid"`
	Hour		int `json:"hour"`
}