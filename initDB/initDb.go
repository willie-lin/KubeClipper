package initDB

import (
	"github.com/jinzhu/gorm"
	"log"
)

var Db *gorm.DB

func init()  {
	var err error


	Db, err = gorm.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/kubec")
	if err != nil {
		log.Panic("err", err.Error())
	}
}