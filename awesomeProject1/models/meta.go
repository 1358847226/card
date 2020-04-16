package models

type 	Meta struct {
	Meta_id int `json:"meta_id"`
	Title string `json:"title"`
	Show bool `json:"show"`
	Icon string `json:"icon"`
	Hidden_header_content bool `json:"hidden_header_content"`
	Keep_alive bool `json:"keep_alive"`
}
