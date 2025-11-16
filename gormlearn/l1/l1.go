package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("===========创建表users==============")
	db.AutoMigrate(&user{})
	db.AutoMigrate(&Dog{})
	db.AutoMigrate(&cat{})

	fmt.Println("===========添加 user dog cat==============")
	db.Debug().Create(&user{Id: 1, Name: "chenhy", Age: 18})

	for i := 0; i < 10; i++ {
		db.Debug().Create(&Dog{uint64(10 + i), "dog" + strconv.Itoa(10+i), 1})
		db.Debug().Create(&cat{uint64(20 + i), "cat" + strconv.Itoa(20+i), 1})
	}

	fmt.Println("===========查询==============")

	var dogs []Dog
	var cats []cat
	result1 := db.Find(&dogs, " id > ?", 15)
	result2 := db.Debug().Find(&cats, "name = ?", "cat23")

	fmt.Println(result1.RowsAffected, result2.RowsAffected)
	fmt.Println(dogs, cats)

}

type Animal interface {
	eat()

	han() string
}

type Dog struct {
	Id     uint64 `grom:"primary key"`
	Name   string
	UserId uint64
}

/*func (d *Dog) eat() {
	fmt.Println(d.name, "is eatting")
}

func (d Dog) han() string {
	return d.name + " is han"
}*/

type cat struct {
	Id     uint64 `grom:"primary key"`
	Name   string
	UserId uint64
}

func (c *cat) eat() {
	fmt.Println("cat ", c.Name, "is eatting")
}

func (c *cat) han() string {
	return "cat " + c.Name + " is han"
}

type user struct {
	Id   uint64 `grom:"primary key"`
	Name string
	Age  uint
}
