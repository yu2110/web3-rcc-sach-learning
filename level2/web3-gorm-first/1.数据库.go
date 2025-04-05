package main

// import (
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	Username string `gorm:"unique"`
// 	Email    string
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

// 	newUser := &User{Username: "sun", Email: "sun@163.com"}
// 	err = createUser(db, newUser)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("created user :", newUser)

// 	//query
// 	userID := newUser.ID
// 	user, err := getUserByID(db, userID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("user by id :", user)

// 	//update user
// 	user.Email = "sun1@163.com"
// 	err = updateUser(db, user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("update user :", user)

// 	//delete user
// 	err = deleteUser(db, user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("delete user :", user)

// }

// func connectToMySQL() (*gorm.DB, error) {
// 	dsn := "root:root1234@(127.0.0.1:3306)/gorm?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// func createUser(db *gorm.DB, user *User) error {
// 	result := db.Create(user)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func getUserByID(db *gorm.DB, userID uint) (*User, error) {
// 	var user User
// 	result := db.First(&user, userID)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &user, nil

// }

// func updateUser(db *gorm.DB, user *User) error {
// 	result := db.Save(user)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }

// func deleteUser(db *gorm.DB, user *User) error {
// 	result := db.Delete(user)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }
