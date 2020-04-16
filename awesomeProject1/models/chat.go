package models

import "time"

type Chat struct {
	Chat_id			int16			`json:"chat_id"`
	Chat_userA_account	string			`json:"chat_user_a_id"`
	Chat_userB_account	string			`json:"chat_user_b_id"`
	Chat_userA_name	string			`json:"chat_user_a_name"`
	Chat_userB_name	string			`json:"chat_user_b_name"`
	Chat_content	string			`json:"chat_content"`
	Chat_date		time.Time		`json:"chat_date"`
}
