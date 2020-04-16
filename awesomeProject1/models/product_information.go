package models


type	Product_information struct {
	Product_information_id			int16		`json:"product_information_id"`
	Product_information_product_id	int16		`json:"product_information_product_id"`
	Product_information_product_name	string		`json:"product_information_product_name"`
	Product_information_type		int16		`json:"product_information_type"`
	Product_information_cont		string		`json:"product_information_cont"`
	Product_information_picture		string		`json:"product_information_picture"`
	Product_information_video		string		`json:"product_information_video"`
	Product_information_pdf			string		`json:"product_information_pdf"`
}