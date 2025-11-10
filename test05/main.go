package main

import (
	"fmt"
)

func main() {
	var a int = 11
	if b := 1; a > 10 {
		b = 2
		// c = 2
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		fmt.Println("其他")
		if c == 3 {
			fmt.Println("c == 3")
		}
		fmt.Println(b)
		fmt.Println(c)
	}

	test(a)
	test02(a)
	test03()
	test04()
}

func test(a int) {
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	default:
		fmt.Println("其他")
	}
}

func test02(a int) {
	v := "a" + "b"
	switch b := 2; a {
	case 1:
		fmt.Println("a == 1", b)
	case 2:
		fmt.Println("a == 2"+string(b), b)
	default:
		fmt.Println("其他"+v, b)
	}
}

func test03() {

	b := 1

	var a interface{}
	a = &b

	switch v := a.(type) {
	case int:
		fmt.Println("int", v)

	case *int:
		fmt.Println("*int", v)
	default:
		fmt.Println("其他", v)
	}
}

type name string
type age int
type salary float64

func constumType(n name, a age, s salary) {
	fmt.Println("name:", n)
	fmt.Println("age:", a)
	fmt.Println("salary:", s)
}

func test04() {
	var n name = "张三"
	switch v := interface{}(n).(type) {
	case name:
		fmt.Println("name", v)
	case age:
		fmt.Println("age", v)
	case salary:
		fmt.Println("salary", v)
	default:
		fmt.Println("其他", v)
	}
}
