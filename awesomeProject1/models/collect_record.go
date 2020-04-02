package models

import "time"

type	Collect_record struct {
	Collect_id			int16		`json:"collect_id"`
	User_account		string `json:"user_account"`
	User_name			string		`json:"user_name"`
	Card_id				int		`json:"card_id"`
	Collect_mark		bool		`grom:"column:collect_mark" json:"collect_mark"`
	Collect_date		time.Time	`grom:"column:collect_date" json:"collect_date"`
}


type Collect_card struct {
	Card_id int `json:"card_id"`
	Collect_date time.Time `json:"collect_date"`
	Username	string `json:"username"`
	Card Card_warehouse `json:"card"`
}

type Collect_card2 struct {
	Card_id int `json:"card_id"`
	Collect_date string `json:"collect_date"`
	Username	string `json:"username"`
	Card Card_warehouse `json:"card"`
}

type Return_card2 struct {
	Card_id int `json:"card_id"`
	Return_date string `json:"collect_date"`
	Username	string `json:"username"`
	Card Card_warehouse `json:"card"`
}