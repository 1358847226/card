package models

type	Card_warehouse struct {
	Card_id					int16			`gorm:"primary_key" json:"card_id"`
	Card_company_id			int			`gorm:"column:card_company_id" json:"card_company_id"`
	Card_user_account		string			`gorm:"column:card_user_account" json:"card_user_account"`
	Card_user_name			string			`gorm:"column:card_user_name" json:"card_user_name"`
	Card_company_name		string			`gorm:"column:card_company_name" json:"card_company_name"`
	Card_address			string			`gorm:"column:card_address" json:"card_address"`
	Card_phone				string			`gorm:"column:card_phone" json:"card_phone"`
	Card_company_post		string			`gorm:"column:card_company_post" json:"card_company_post"`
	Card_email				string			`gorm:"column:card_email" json:"card_email"`
	Card_introduction		string			`gorm:"column:card_introduction" json:"card_introduction"`
	Card_picture			string			`gorm:"column:card_picture" json:"card_picture"`
	Card_video				string			`gorm:"column:card_video" json:"card_video"`
	Card_privacy			bool			`gorm:"column:card_privacy" json:"card_privacy"`
	Card_browse				int16			`gorm:"column:card_browse" json:"card_browse"`
	Card_collect			int16			`gorm:"column:card_collect" json:"card_collect"`
	Website_address			string			`gorm:"column:website_address" json:"website_address"`
	Card_mould_id			int				`gorm:"column:card_mould_id" json:"card_mould_id"`
	Mark					bool			`gorm:"column:mark" json:"mark"`
	Web_introduction		string			`json:"web_introduction"`
	Web_picture				string			`json:"web_picture"`
	Web_video				string			`json:"web_video"`
	Web_introduction_public		string			`json:"web_introduction_public"`
	Web_picture_public			string			`json:"web_picture_public"`
	Web_video_public			string			`json:"web_video_public"`
	Keyword						string			`json:"keyword"`
	Avatar						string			`json:"avatar"`
}
type Card struct {
	Card_id					int16			`gorm:"primary_key" json:"card_id"`
	Card_company_id			int			`gorm:"column:card_company_id" json:"card_company_id"`
	Card_user_account		string			`gorm:"column:card_user_account" json:"card_user_account"`
	Card_user_name			string			`gorm:"column:card_user_name" json:"card_user_name"`
	Card_company_name		string			`gorm:"column:card_company_name" json:"card_company_name"`
	Card_address			string			`gorm:"column:card_address" json:"card_address"`
	Card_phone				string			`gorm:"column:card_phone" json:"card_phone"`
	Card_company_post		string			`gorm:"column:card_company_post" json:"card_company_post"`
	Card_email				string			`gorm:"column:card_email" json:"card_email"`
	Card_introduction		string			`gorm:"column:card_introduction" json:"card_introduction"`
	Card_picture			[]string			`gorm:"column:card_picture" json:"card_picture"`
	Card_video				[]string			`gorm:"column:card_video" json:"card_video"`
	Card_privacy			bool			`gorm:"column:card_privacy" json:"card_privacy"`
	Card_browse				int16			`gorm:"column:card_browse" json:"card_browse"`
	Card_collect			int16			`gorm:"column:card_collect" json:"card_collect"`
	Website_address			string			`gorm:"column:website_address" json:"website_address"`
	Card_mould_id			int				`gorm:"column:card_mould_id" json:"card_mould_id"`
	Mark					bool			`gorm:"column:mark" json:"mark"`
	Web_introduction		string			`json:"web_introduction"`
	Web_picture				[]string			`json:"web_picture"`
	Web_video				[]string			`json:"web_video"`
	Web_introduction_public		[]string			`json:"web_introduction_public"`
	Web_picture_public			[]string			`json:"web_picture_public"`
	Web_video_public			[]string			`json:"web_video_public"`
	Keyword						[]string			`json:"keyword"`
	Avatar						string			`json:"avatar"`
}