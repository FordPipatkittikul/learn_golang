package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Sqllogger struct {
	logger.Interface
}

func (l Sqllogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n============================\n", sql)
}

var db *gorm.DB

func main() {
	dsn := "root:admin1234@tcp(127.0.0.1:3306)/godb?parseTime=true"
	dial := mysql.Open(dsn)

	var err error
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &Sqllogger{},
		DryRun: false, // only see log don't actually doing it
	})
	if err != nil{
		panic(err)
	}


	// err = db.Migrator().CreateTable(Gender{})

	// err = db.AutoMigrate(Gender{}, Test{})
	// if err != nil {
	// 	fmt.Printf("AutoMigrate error: %v\n", err)
	// }

	CreateGender("Female")

}

type Gender struct {
	ID   uint
	Name string `gorm:"unique;size(10)"`
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

// func GetGenders() {
// 	db.
// }

type Test struct {
	gorm.Model
	Code uint   `gorm:"comment:This is Code"`
	Name string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
}

func (t Test) TableName() string {
	return "MyTest"
}