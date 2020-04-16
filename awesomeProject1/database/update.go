package database

import (
	"awesomeProject1/connect"
	_ "awesomeProject1/connect"
	"awesomeProject1/models"
	_ "fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func Update_user(data   models.User, id int16) {
			var user  models.User
			connect.Db.Debug().Model(&user).Where("user_id = ?",id).Updates(&data)
}

func Update_user_name(account string, name string) {
	var user  models.User
	connect.Db.Debug().Model(&user).Where("account = ?",account).Update("user_name",name)
}
func Update_user_avatar(account string, avatar string) {
	var user  models.User
	connect.Db.Debug().Model(&user).Where("account = ?",account).Update("avatar",avatar)
}


func Update_collectf(userid string, cardid int) error {
	var collect  models.Collect_record
	return connect.Db.Debug().Model(&collect).Where("user_account = ? and card_id = ?",userid,cardid).Updates(map[string]interface{}{"collect_mark":false,"collect_date":time.Now()}).Error
}

func Update_collectt(userid string, cardid int) error {
	var collect  models.Collect_record
	return connect.Db.Debug().Model(&collect).Where("user_account = ? and card_id = ?",userid,cardid).Updates(map[string]interface{}{"collect_mark":true,"collect_date":time.Now()}).Error
}

func Updata_card(cardid int,data models.Card_warehouse){
	var card models.Card_warehouse
	connect.Db.Debug().Model(&card).Where("card_id = ?",cardid).Update(data)
	connect.Db.Debug().Model(&card).Where("card_id = ?",cardid).Updates(map[string]interface{}{"card_picture": data.Card_picture, "card_video":data.Card_video})
}

func Updata_card_avatar(cardid int,avatar string){
	var card models.Card_warehouse
	connect.Db.Debug().Model(&card).Where("card_id = ?",cardid).Update("avatar",avatar)
}

func Updata_card_mark(account string)  {
	var card models.Card_warehouse
	connect.Db.Model(&card).Where("card_user_account = ?",account).Update("mark",0)
}

func Updata_card_mark0(account string) error {
	var card models.Card_warehouse
	return connect.Db.Model(&card).Where("card_user_account = ?",account).Update("mark",0).Error
}

func Updata_card_mark1(cardid int) error {
	var card models.Card_warehouse
	return connect.Db.Debug().Model(&card).Where("card_id = ?",cardid).Update("mark",true).Error
}

func Updata_browse(card_id int16) {
	var browse  models.Card_warehouse
	connect.Db.Model(&browse).Where("card_id = ?",card_id).Update("card_browse",gorm.Expr("card_browse + 1"))
}

func Collect_add(card_id int) {
	var collect  models.Card_warehouse
	connect.Db.Debug().Model(&collect).Where("card_id = ?",card_id).Update("card_collect",gorm.Expr("card_collect + 1"))
}

func Browse_add(card_id int) {
	var collect  models.Card_warehouse
	log.Println("567567")
	connect.Db.Debug().Model(&collect).Where("card_id = ?",card_id).Update("card_browse",gorm.Expr("card_browse + 1"))
}

func Collect_minus(card_id int) {
	var collect  models.Card_warehouse
	connect.Db.Model(&collect).Where("card_id = ?",card_id).Update("card_collect",gorm.Expr("card_collect - 1"))
}

func Updata_user(account string, qty int)  {
	var user models.User
	connect.Db.Model(&user).Where("account = ?",account).Update("card_quantity",qty)
}
func Updata_user_name(account string, name string, avatar string)  {
	var user models.User
	connect.Db.Model(&user).Where("account = ?",account).Update("user_name",name)
	connect.Db.Model(&user).Where("account = ?",account).Update("avatar",avatar)
}

func Updata_password(admin string, password string) error {
	var user models.User
	password = MD5(password)
	log.Printf(admin)
	log.Println(password)
	return  connect.Db.Model(&user).Where("admin_id = ?",admin).Update("user_password",password).Error
}

func Add_search_note(account string, keyword string) {
	var note models.Search_notes
	connect.Db.Model(&note).Where("account = ? and keyword = ?",account,keyword).Update("num",gorm.Expr("num + 1"))
}

func Add_search_num( keyword string) {
	var num models.Search_keyword_num
	connect.Db.Model(&num).Where("keyword = ?",keyword).Update("num",gorm.Expr("num + 1"))
}

func Add_search_user_num( keyword string) {
	var num models.Search_keyword_num
	connect.Db.Model(&num).Where("keyword = ?",keyword).Update("user_num",gorm.Expr("user_num + 1"))
}

func Updata_share(account string) {
	var share models.User
	connect.Db.Model(&share).Where("account = ?",account).Update("user_share",gorm.Expr("user_share + 1"))
}

