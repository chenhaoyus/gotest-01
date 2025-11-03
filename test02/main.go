package main

import (
	"errors"
	"fmt"
)

func main() {
	var f1 float32 = 10.0

	var f2 = 12.3

	f1 = float32(f2)

	fmt.Println("f1:", f1)
	fmt.Println("f2:", f2)

	var a string = "hello 你好"

	var bytes []byte = []byte(a)
	fmt.Println("bytes:", bytes)

	var str string = string(bytes)
	fmt.Println("str:", str)

	var unit8s []uint8 = []uint8(a)

	fmt.Println("unit8s:", unit8s)

	var runes []rune = []rune(a)
	fmt.Println("runes:", runes)

	fmt.Println("runes len:", len(runes))
	fmt.Println("string len:", len(a))

	fmt.Println("aaa:", g)
	fmt.Println("cccc:", f)

	fmt.Println("===============================================================")

	var p1 *int
	var p2 *string

	i := 1
	s := "Hello"
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	p1 = &i
	p2 = &s

	p3 := &p2
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(*p3)

	var o1 = bbg{1, "da"}

	var z1 *bbg = &o1

	fmt.Println("z1:", z1)
	fmt.Println("z1.a1:", z1.a1)
	fmt.Println("z1.b1:", z1.b1)

	fmt.Println("===============================================================")

	var d interface{} = "hello"

	str2, ok := d.(string)
	if ok {
		fmt.Println("str2:", str2)
	} else {
		fmt.Println("d is not a string")
	}

	v, r := d.(int)

	fmt.Println("v:", v)
	fmt.Println("r:", r)

	fmt.Println("===============================================================")

	err := errors.New("this is an error")

	fmt.Println("err:", err.Error())

}

type bbg struct {
	a1 int
	b1 string
}

type s1 struct {
	a1 int
	b1 string
}

var g, f = dsada()

var (
	i = 1
	b = "da"
)

func dsada() (aaa int, ccc string) {
	return i, b
}
