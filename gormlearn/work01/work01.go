package main

import (
	"fmt"

	"github.com/go-test01/gormlearn/globa"

	"gorm.io/gorm"
)

func main() {
	db := globa.DB
	db.AutoMigrate(&Student{})
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"

	student := &Student{ID: 1, Name: "张三", Age: 20, Grade: "三年级"}
	insert(student, db)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息

	students := []Student{}
	query(&students, "age > 18", db)
	fmt.Println(students)

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	updatecount := update(&Student{}, "name = 张三", "grade", "四年级", db)
	fmt.Println(updatecount)

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录
	delete("age > 15", db)
}

func insert(student *Student, db *gorm.DB) {

	db.Create(&student)
}

func query(students *[]Student, where string, db *gorm.DB) []Student {

	db.Debug().Find(students, where)

	return *students
}

func update(student *Student, where string, column string, update string, db *gorm.DB) int64 {
	dd := db.Debug().Model(student).Where(where).Update(column, update)
	return dd.RowsAffected
}

func delete(where string, db *gorm.DB) int64 {
	dd := db.Debug().Where(where).Delete(&Student{})
	return dd.RowsAffected
}

type Student struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Age   int
	Grade string
}
