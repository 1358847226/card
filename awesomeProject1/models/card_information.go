package models


type Card_information struct {
	Card_information_id			int16			`json:"card_information_id"`
	Card_information_type		int8			`json:"card_information_type"`
	Card_information_card_id	int16			`json:"card_information_card_id"`
	Card_information_cont		string			`json:"card_information_cont"`
	Card_information_picture	string			`json:"card_information_picture"`
	Card_information_video		string			`json:"card_information_video"`
	Card_information_pdf		string			`json:"card_information_pdf"`

}