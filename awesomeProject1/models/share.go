package models

import "time"

type Share struct {
	Card_id int `json:"card_id"`
	Share_account string `json:"share_account"`
	Share_name string `json:"share_name"`
	Visit_account string `json:"visit_account"`
	Visit_name string `json:"visit_name"`
	Time time.Time `json:"time"`
}

type Share_2 struct {
	Card_id int `json:"card_id"`
	Share_account string `json:"share_account"`
	Share_name string `json:"share_name"`
	Visit_account string `json:"visit_account"`
	Visit_name string `json:"visit_name"`
	Time string `json:"time"`
	Avatar string `json:"avatar"`
}
