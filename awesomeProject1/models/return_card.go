package models

import "time"

type Return_card struct {
	ReturnId int `json:"return_id"`
	ReturnAccount string `json:"return_account"`
	AcceptAccount string `json:"accept_account"`
	CardId int `json:"card_id"`
	Date time.Time `json:"date"`
}
