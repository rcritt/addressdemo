package apis

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rcritt/addressdemo/orm"
	"os"
)

// Keep this open as we will keep needing it.
var db, _ = gorm.Open("mysql", os.Getenv("DB_URL"))

func init() {
	db.DB()
}

func Create(newAddressInfo orm.AddressInfo) {
	db.Create(&newAddressInfo)
}

func List() []orm.AddressInfo {
	var addresses []orm.AddressInfo
	db.Find(&addresses)

	return addresses
}
