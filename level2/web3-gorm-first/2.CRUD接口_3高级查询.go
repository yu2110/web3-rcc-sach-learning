package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// )

// type User struct {
// 	// ID       uint   `gorm:"primaryKey"`
// 	gorm.Model
// 	Name     string `gorm:"unique"`
// 	Age      int
// 	Birthday time.Time
// }

// type APIUser struct {
// 	ID   uint
// 	Name string
// }

// func main() {

// 	db, err := connectToMySQL()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// defer db.Close()
// 	// Perform database migration
// 	err = db.AutoMigrate(&User{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// fmt.Println("db :", db)
// 	// Your CRUD operations go here
// 	// //创建记录
// 	// users := []*User{
// 	// 	{Name: "zhang", Age: 20, Birthday: time.Now()},
// 	// 	{Name: "sun", Age: 19, Birthday: time.Now()},
// 	// }

// 	// user := &User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
// 	// result := db.Create(user)
// 	// fmt.Println(result)

// 	//智能选择字段
// 	// var user User
// 	// var apiusers []APIUser
// 	// // 在查询时，GORM 会自动选择 `id `, `name` 字段
// 	// db.Model(&User{}).Limit(10).Find(&apiusers)
// 	// // SQL: SELECT `id`, `name` FROM `users` LIMIT 10

// 	//锁
// 	var users []User
// 	// 基本的 FOR UPDATE 锁
// 	db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)
// 	// SQL: SELECT * FROM `users` FOR UPDATE

// 	fmt.Println("db user name ", users)

// }

// func connectToMySQL() (*gorm.DB, error) {
// 	dsn := "root:root1234@(127.0.0.1:3306)/gorm?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
