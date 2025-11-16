package main

import (
	"fmt"
	"github.com/gormlearn/globa"
	"github.com/jmoiron/sqlx"
)

func init() {
	db := globa.DB

	db.AutoMigrate(&Employee{})

	employees := []Employee{
		{Name: "张三", Department: "技术部", Salary: 15000},
		{Name: "李四", Department: "技术部", Salary: 18000},
		{Name: "王五", Department: "销售部", Salary: 12000},
		{Name: "赵六", Department: "技术部", Salary: 20000},
		{Name: "钱七", Department: "人事部", Salary: 10000},
		{Name: "孙八", Department: "技术部", Salary: 22000},
		{Name: "周九", Department: "财务部", Salary: 13000},
	}

	db.Create(employees)

	db.AutoMigrate(&Book{})
	bookTemplates := []Book{
		{Title: "Go语言实战", Author: "Alan A. Donovan", Price: 89.00},
		{Title: "深入理解计算机系统", Author: "Randal E. Bryant", Price: 139.00},
		{Title: "算法导论", Author: "Thomas H. Cormen", Price: 128.00},
		{Title: "设计模式", Author: "Erich Gamma", Price: 79.00},
		{Title: "Clean Code", Author: "Robert C. Martin", Price: 69.00},
		{Title: "重构", Author: "Martin Fowler", Price: 99.00},
		{Title: "HTTP权威指南", Author: "David Gourley", Price: 118.00},
		{Title: "数据库系统概念", Author: "Abraham Silberschatz", Price: 149.00},
	}

	db.Create(bookTemplates)
}

func main() {
	test01()
	test02()

}

func test01() {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=true")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
	var emp []Employee = []Employee{}
	err2 := db.Select(&emp, "select * from employees where department = ?", "技术部")

	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(emp)

	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
	var emp2 []Employee = []Employee{}
	err3 := db.Select(&emp2, "select * from employees order by salary desc limit 1")

	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(emp2)
}

func test02() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=true"

	connect, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer connect.Close()

	book := []Book{}
	err3 := connect.Select(&book, "select * from books where price > ?", 50)

	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(book)
}

type Book struct {
	ID     int `gorm:"primary_key;auto_increment"`
	Title  string
	Author string
	Price  int
}

type Employee struct {
	ID         int `gorm:"primary_key:auto_increment"`
	Name       string
	Department string
	Salary     int
}
