package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//panic(err)
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	error := db.AutoMigrate(&user)

	if nil != error {
		panic(error)
	}

	result := db.Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	var user1 []User
	//db.Debug().First(&user1, 1)
	db.Debug().Find(&user1, "id > 1 and id < 4")

	fmt.Println(user1)
	var user2 User
	db.Debug().Unscoped().Model(&user2).Where("id > 1 and deleted_at is not null").Update("deleted_at", time.Now())
}

type User struct {
	gorm.Model
	Name       string
	Age        int
	Birthday   time.Time
	creditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
