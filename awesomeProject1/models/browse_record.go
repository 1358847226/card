package models

import "time"

type BrowseRecord struct {

	Useraccount   	string		`gorm:"column:user_account" json:"useraccount"`
	Username		string		`gorm:"column:user_name" json:"username"`
	Cardid			int16		`gorm:"column:card_id" json:"cardid"`
	Date			time.Time	`gorm:"column:browse_date" json:"date"`
}
