package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin1234@tcp(127.0.0.1:3306)/godb?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil{
		panic(err)
	}
	db.Migrator().CreateTable(Gender{})
}

type Gender struct {
	ID 		uint
	Name 	string
}