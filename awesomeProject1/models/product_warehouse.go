package models

import "time"

type Product_warehouse struct {
	Product_id			int16		`json:"product_id"`
	Product_company_id	int16		`json:"product_company_id"`
	Product_company_name	string		`json:"product_company_id"`
	Product_name		string		`json:"product_name"`
	Product_picture		string		`json:"product_picture"`
	Product_keyword		string		`json:"product_keyword"`
	Product_time		time.Time	`json:"product_time"`
}
