package database

import (
	"awesomeProject1/connect"
	"awesomeProject1/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func Createuser(data models.User) *gorm.DB {

	var user  models.User
	user = data
	err := connect.Db.Create(&user)
	return err

}


func Createcard(data models.Card_warehouse) *gorm.DB {
	var card  models.Card_warehouse
	card = data

		var err = connect.Db.Create(&card)
		return  err

}

func Createbrowse(account string,username string ,card_id int) error {

	var browse models.BrowseRecord
	browse = models.BrowseRecord{
		Useraccount: account,
		Cardid:      int16(card_id),
		Username: 	username,
		Date:        time.Now(),
	}
	return connect.Db.Create(&browse).Error
}

func Createinformation(data models.Card_information) error {
	var information models.Card_information
	information = data
	return connect.Db.Create(&information).Error
}

func Createchat(data models.Chat) error {
	var chat models.Chat
	chat = data
	return connect.Db.Create(&chat).Error
}

func	Createcollect(data models.Collect_record) *gorm.DB {
	var collect models.Collect_record
	collect = data
	return connect.Db.Create(&collect)
}

func Createcompany(data models.Company_warehouse)  {
	company := data
	err := connect.Db.Create(&company)
	if err !=nil{
		fmt.Println("创建失败，公司名为空或已存在！")
	}
}

func Createkeyword(data models.Keyword_warehouse) bool {
	keyword := data
	if (len(Search_keyword_bycon(keyword.Keyword_content))>0){
		return false
	}else
	{
		if connect.Db.Create(&keyword).Error != nil{
			return false
		}else{
			return true
		}
	}


}

func Createneed(data models.Need_warehouse) error {
	need := data
	return connect.Db.Create(&need).Error
}

func Create_cardmould(data models.Card_mould) error {
	mould := data
	return connect.Db.Debug().Create(&mould).Error

}

func Create_loginnotes(account string) error {
	notes := models.Login_notes{
		Account: account,
		Time:    time.Now(),
		Hour:	time.Now().Hour(),
	}
	return connect.Db.Debug().Create(&notes).Error
}

func Create_search_notes(data models.Search_notes)  {
	note := data
	connect.Db.Create(&note)
}

func Create_search_num(keyword string)  {
	num := models.Search_keyword_num{
		Keyword:  keyword,
		Num:      1,
		User_num: 1,
	}
	connect.Db.Create(&num)
}

func Create_share(card_id int, share_account string, visit_account string,sname string,vname string) error {
	share := models.Share{
		Card_id:       card_id,
		Share_account: share_account,
		Visit_account: visit_account,
		Share_name:    sname,
		Visit_name:  	vname,
		Time:          time.Now(),
	}
	return connect.Db.Create(&share).Error
}

func CreateReturnCard(data models.Return_card)  {
	connect.Db.Create(&data)
}
