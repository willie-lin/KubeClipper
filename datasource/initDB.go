package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init()  {

	var  err error

	Db, err = gorm.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/cdn")

	if err != nil {
		log.Panicln("err:", err.Error())
	}
}