package model

import (
	"KubeClipper/datasource"
	"github.com/jinzhu/gorm"
)

type Cdnip struct {
	gorm.Model
	Ip	string		`json:"ip_addr"`

}

func init()  {
	table := datasource.Db.HasTable(Cdnip{})
	if !table {
		datasource.Db.CreateTable(Cdnip{})
	}
}

func (cndip Cdnip) QueryByIp() Cdnip {
	datasource.Db.First(&cndip, cndip.Ip)
	return cndip
}