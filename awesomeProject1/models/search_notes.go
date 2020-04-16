package models

import "time"

type Search_notes struct {
	Account string `json:"account"`
	Keyword string `json:"keyword"`
	Num		int `json:"num"`
	Time time.Time	`json:"time"`
	Name string `json:"name"`
	First_search bool `json:"first_search"`
}
