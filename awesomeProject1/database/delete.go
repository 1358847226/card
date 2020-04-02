package database

import (
	"awesomeProject1/connect"
	_ "awesomeProject1/connect"
	"awesomeProject1/models"
	_ "fmt"
	"log"
	"os"
	"strings"
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

func	Delete_uploadFile(file []string) {
	for i := 0 ; i < len(file) ; i++{
		comma := strings.Index(file[i], "u")
		localFile := "./" + file[i][comma:]
		err := os.Remove(localFile)
		log.Println(localFile)
		log.Println(err)
	}
}
