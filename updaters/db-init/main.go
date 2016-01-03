package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rcritt/addressdemo/orm"
	"os"
)

func main() {
	fmt.Println("First time initialization of the database...")

	// Connect to the database.
	db, _ := gorm.Open("mysql", os.Getenv("DB_URL"))

	db.DB()
	// Create our table.
	db.CreateTable(&orm.AddressInfo{})

	// Report on the creation & exit.
	fmt.Println("Database initialization has concluded.")
}
