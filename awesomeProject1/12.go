package main


import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "fmt"
)

type Users struct {
	id	int `gorm:"type:int(11);column:user_id"`
	name string `gorm:"type:varchar(255);column:user_name"`
	password string `gorm:"type:varchar(255);column:user_password"`
	power int `gorm:"type:int(4);column:user_power"`
	share int `gorm:"type:int(11);column:user_share"`
}



var db *gorm.DB
func main() {
	var err error

	db, err = gorm.Open("mysql", "root:123456@/card?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	/*user := &Users{
		id:2,
		name:"wang",
		password:"123456",
		power:1,
		share:0,
	}
	fmt.Print(user)
	db.Table("user").Create(user)*/
	user := &Users{
		id:       1,
		name:     "",
		password: "",
		power:    0,
		share:    0,
	}
	db.Model(&user).Update("user_id","7")
}

