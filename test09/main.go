package main

import "fmt"

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
}
