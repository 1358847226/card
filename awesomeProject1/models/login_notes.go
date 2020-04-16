package models

import "time"

type Login_notes struct {
	Account string `json:"account"`
	Time time.Time	`json:"time"`
	Hour int		`json:"hour"`
}


type User_visit struct {
	All_visit int `json:"all_visit"`
	Today_visit int `json:"today_visit"`
	Yestoday_visit int `json:"yestoday_visit"`
	Rate string `json:"rate"`
}