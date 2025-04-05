package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// )

// type CreditCard struct {
// 	gorm.Model
// 	Number string
// 	UserID uint
// }

// type User struct {
// 	gorm.Model
// 	Name       string `gorm:"unique""default:root"`
// 	CreditCard CreditCard
// }

// func main() {

// 	db, err := connectToMySQL()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// defer db.Close()
// 	// Perform database migration
// 	err = db.AutoMigrate(&CreditCard{})
// 	err = db.AutoMigrate(&User{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("db :", db)
// 	// Your CRUD operations go here
// 	user := &User{
// 		Name:       "jinzhu",
// 		CreditCard: CreditCard{Number: "411111111111"},
// 	}
// 	// db.Create(user)
// 	//Select, Omit方法来跳过关联更新
// 	// db.Omit("CreditCard").Create(&user)

// 	// skip all associations
// 	// db.Omit(clause.Associations).Create(&user)

// 	//Upsert及冲突
// 	// Do nothing on conflict
// 	// db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

// 	// Update columns to default value on `id` conflict
// 	db.Clauses(clause.OnConflict{
// 		Columns:   []clause.Column{{Name: "name"}},
// 		DoUpdates: clause.Assignments(map[string]interface{}{"updated_at": time.Now()}),
// 	}).Create(&user)
// 	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET ***; SQL Server
// 	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE ***; MySQL
// 	//

// }

// func connectToMySQL() (*gorm.DB, error) {
// 	dsn := "root:root1234@(127.0.0.1:3306)/gorm?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
