package models

import "time"

type Need_warehouse struct {
	Need_id					int16		`json:"need_id"`
	Need_user_id			int16		`json:"nedd_user_id"`
	Need_user_name			string		`json:"need_name"`
	Need_name				string		`json:"need_name"`
	Need_content			string		`json:"need_content"`
	Need_keyword			string		`json:"need_keyword"`
	Need_starttime			time.Time	`json:"need_startime"`
	Need_endtime			time.Time	`json:"need_endtime"`
}
