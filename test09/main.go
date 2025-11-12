package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=======类型转换==========")

	var a int = 10
	var b byte = 5
	var c float32 = 3.14

	fmt.Println("a:", a, "b:", b, "c:", c)

	a = int(b)
	fmt.Println("a:", a)

	b = byte(a)
	fmt.Println("b:", b)

	test01()
	test02()
	test03()
}

func test01() {
	var a string = "hello"

	var b []byte = []byte(a)

	fmt.Println(b)

	var c string = "你好"

	var d []rune = []rune(c)
	var e []byte = []byte(c)

	fmt.Println(d)
	fmt.Println(e)

	var f = string(e)
	var l = string(d)

	fmt.Println(f)
	fmt.Println(l)
}

func test02() {

	a := "hello 你好！"

	for i := 0; i < len(a); i++ {
		fmt.Printf("第%d个字节：%x\n", i, a[i])
	}

	for i, v := range a {
		fmt.Printf("第%d个字符：%c\n", i, v)
	}
}

func test03() {
	a, err := strconv.Atoi("20")

	if err != nil {
		panic(err)
	}

	fmt.Println(a)

	b := strconv.Itoa(a)

	fmt.Println(b)
}
