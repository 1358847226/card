package database

import (
	"awesomeProject1/connect"
	_ "awesomeProject1/connect"
	"awesomeProject1/models"
	_ "fmt"
)

func	Delete(id int16){
	user := models.User{}
	connect.Db.Where("user_id = ?",id).Delete(&user)

}

func	Admin_Deletecard_bycardid(id int) error {
	card := models.Card_warehouse{}
	err := connect.Db.Where("card_id = ?",id).Delete(&card).Error
	return err
}

func	Deletecard_bycardid(cardid int,userid string) error {
	card := models.Card_warehouse{}
	err := connect.Db.Where("card_id = ? and card_user_account = ?",cardid,userid).Delete(&card).Error
	return err
}