package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	for i := 0; i < 10; i++ {
		fmt.Println("======", i, "==========")
	}

	// 方式2
	b := 1
	for b < 10 {
		fmt.Println("方式2，第", b, "次循环")
		b++
	}

	test01()
	test02()

	a := aaa{name: "张三", age: 18}

	functest = a.test04

	functest()
}

func test01() {
	clousre := func() {
		fmt.Println("this is a clousre function")
	}
	clousre()
}

func test02() {

	clousre := test03()

	clousre()
}

func test03() func() {

	return func() {
		fmt.Println("this is a clousre function")
	}
}

var functest func()

type aaa struct {
	name string
	age  int
}

func (a *aaa) test04() {
	fmt.Println("name:", a.name, "age:", a.age)
}
