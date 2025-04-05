package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

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
// 	// 检索单个对象
// 	// 获取第一条记录（主键升序）
// 	var user User
// 	// db.First(&user)
// 	// SELECT * FROM users ORDER BY id LIMIT 1;

// 	// 获取一条记录，没有指定排序字段
// 	//db.Take(&user)
// 	// SELECT * FROM users LIMIT 1;

// 	// 获取最后一条记录（主键降序）
// 	// db.Last(&user)
// 	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// 	// 根据主键检索
// 	// db.First(&user, 2)
// 	// SELECT * FROM users WHERE id = 10;

// 	// String 条件
// 	// Get first matched record
// 	db.Where("name = ?", "jinzhu").First(&user)
// 	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

// 	// Get all matched records
// 	// db.Where("name <> ?", "jinzhu").Find(&users)
// 	// SELECT * FROM users WHERE name <> 'jinzhu';

// 	//Struct & Map 条件

// 	// Struct
// 	db.Where(&User{Name: "jinzhu", Age: 18}).First(&user)
// 	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

// 	// Map
// 	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
// 	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

// 	fmt.Println("db user name ", user.Name)

// }

// func connectToMySQL() (*gorm.DB, error) {
// 	dsn := "root:root1234@(127.0.0.1:3306)/gorm?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
