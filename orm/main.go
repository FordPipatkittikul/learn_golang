package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	// err = db.AutoMigrate(Gender{}, Test{}, Customer{})
	// if err != nil {
	// 	fmt.Printf("AutoMigrate error: %v\n", err)
	// }

	// CreateGender("Female")
	// GetGenders()
	// GetGender(1)
	// GetGenderByName("Female")
	// UpdateGender(1, "Gay")
	// UpdateGender2(2, "Lesbian")
	// DeleteGender(1)

	// CreateTest(0, "Test1")
	// CreateTest(0, "Test2")
	// CreateTest(0, "Test3")
	// SoftDeleteTest(3)
	// HardDeleteTest(2)

	// db.Migrator().CreateTable(Customer{})
	// CreateCustomer("joke", 2)
	GetCustomers()
}

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender
	GenderID uint
}

func GetCustomers() {
	customers := []Customer{}
	tx := db.Preload(clause.Associations).Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, customer := range customers {
		fmt.Printf("%v|%v|%v\n", customer.ID, customer.Name, customer.Gender.Name)
	}
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{Name: name, GenderID: genderID}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
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

func GetGenders() {
	genders := []Gender{}
	tx := db.Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, "id=?",id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenderByName(name string) {
	gender := Gender{}
	tx := db.First(&gender, "name=?", name)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func UpdateGender(id uint, name string) {
	gender := Gender{}
	tx := db.First(&gender, "id=?",id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}

func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}

func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
}

type Test struct {
	gorm.Model
	Code uint   `gorm:"comment:This is Code"`
	Name string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
}

func (t Test) TableName() string {
	return "MyTest"
}

func CreateTest(code uint, name string){
	test := Test{Code: code, Name: name}
	db.Create(&test)
}

func GetTests() {
	tests := []Test{}
	db.Find(&tests)
}


func SoftDeleteTest(id uint) {
	db.Delete(&Test{}, id)
}

func HardDeleteTest(id uint) {
	db.Unscoped().Delete(&Test{}, id)
}