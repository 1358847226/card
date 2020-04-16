package models

type Card_mould struct {
	Card_mould_id				int16			`json:"card_mould_id"`
	User_account				string			`json:"user_account"`
	Card_mould_style			int16			`json:"card_mould_style"`
	Card_mould_name				string			`json:"card_mould_name"`
	Card_mould_company_id		int16 			`json:"card_mould_company_id"`
	Card_mould_company_name		string			`json:"card_mould_company_name"`
	Card_mould_address			string			`json:"card_mould_address"`
	Card_mould_post				string			`json:"card_mould_post"`
	Card_mould_website_style	int16			`json:"card_mould_website_style"`
	Card_mould_website_address	string			`json:"card_mould_website_address"`
	Web_introduction			string			`json:"web_introduction"`
	Web_picture					string			`json:"web_picture"`
	Web_video					string			`json:"web_video"`
	Web_introduction_public		string			`json:"web_introduction_public"`
	Web_picture_public			string			`json:"web_picture_public"`
	Web_video_public			string			`json:"web_video_public"`
}
