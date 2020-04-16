package models

type website_information struct {
	website_information_id			int16		`json:"website_information_id"`
	website_information_company_id	int16		`json:"website_information_company_id"`
	website_information_user_id		int16		`json:"website_information_user_id"`
	website_information_type		string		`json:"website_information_type"`
	website_information_cont		string		`json:"website_information_cont"`
	website_information_picture		string		`json:"website_information_picture"`
	website_information_video		string		`json:"website_information_video"`
	website_information_pdf			string		`json:"website_information_pdf"`
}