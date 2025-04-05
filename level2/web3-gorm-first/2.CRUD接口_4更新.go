package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	// ID       uint   `gorm:"primaryKey"`
	gorm.Model
	Name     string `gorm:"unique"`
	Age      int
	Birthday time.Time
}

func main() {

	db, err := connectToMySQL()
	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()
	// Perform database migration
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("db :", db)
	// Your CRUD operations go here
	// //创建记录
	users := []*User{
		{Name: "zhang", Age: 20, Birthday: time.Now()},
		{Name: "sun", Age: 19, Birthday: time.Now()},
	}

	// user := &User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := db.Create(users)
	fmt.Println(result)

	//保存所有字段
	// db.Save(&user)

	// fmt.Println("db user name ", users)

}

func connectToMySQL() (*gorm.DB, error) {
	dsn := "root:root1234@(127.0.0.1:3306)/gorm?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
