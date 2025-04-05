package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/google/uuid"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	// ID       uint   `gorm:"primaryKey"`
// 	gorm.Model
// 	Name     string `gorm:"unique"`
// 	Age      int
// 	Birthday time.Time
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

// 	fmt.Println("db :", db)
// 	// Your CRUD operations go here

// 	//创建记录
// 	// users := []*User{
// 	// 	{Name: "zhang", Age: 20, Birthday: time.Now()},
// 	// 	{Name: "sun", Age: 19, Birthday: time.Now()},
// 	// }
// 	// result := db.Create(users)

// 	// 用指定的字段创建记录
// 	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
// 	// result := db.Select("Name", "Age", "CreatedAt").Create(&user)
// 	// // INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
// 	// fmt.Println("result error %v", result.Error)
// 	// fmt.Println("row affected %v", result.RowsAffected)

// 	//给 ‘Omit’ 的字段值
// 	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
// 	// db.Omit("Name", "Age", "CreatedAt").Create(&user)
// 	// // INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

// 	// 批量插入
// 	// var users = []User{{Name: "jinzhu1", Age: 18, Birthday: time.Now()},
// 	// 	{Name: "jinzhu2", Age: 18, Birthday: time.Now()},
// 	// 	{Name: "jinzhu3", Age: 18, Birthday: time.Now()}}
// 	// db.Create(&users)

// 	// batch size 100
// 	// var users = []User{{Name: "jinzhu1", Age: 18, Birthday: time.Now()},
// 	// 	{Name: "jinzhu2", Age: 18, Birthday: time.Now()},
// 	// 	{Name: "jinzhu3", Age: 18, Birthday: time.Now()}}
// 	// db.CreateInBatches(users, 2)

// 	//批量插入 植入钩子
// 	// db = db.Session(&gorm.Session{CreateBatchSize: 2})
// 	// db = db.Session(&gorm.Session{SkipHooks: true}) //跳过钩子
// 	// var users = []User{{Name: "jinzhu1", Age: 18, Birthday: time.Now()},
// 	// 	{Name: "jinzhu2", Age: 18, Birthday: time.Now()},
// 	// 	{Name: "jinzhu3", Age: 18, Birthday: time.Now()}}
// 	// db.Create(&users)

// 	//根据 Map 创建
// 	db.Model(&User{}).Create(map[string]interface{}{
// 		"Name": "jinzhu2", "Age": 18, "Birthday": time.Now(),
// 	})

// }

// func connectToMySQL() (*gorm.DB, error) {
// 	dsn := "root:root1234@(127.0.0.1:3306)/gorm?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// // 创建钩子
// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	uuid := uuid.New()
// 	log.Println("BeforeCreate  uuid===", uuid)

// 	return
// }
