package connect

import (
	"fmt"
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init(){
	var err error

	Db, err = gorm.Open("mysql", "root:123456@/card?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	Db.SingularTable(true)
}


//func Connect() {
//	var err error
//
//	Db, err = gorm.Open("mysql", "root:123456@/card?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(err)
//}
