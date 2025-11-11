package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(a); i++ {
		fmt.Println("======", a[i], "==========")
	}

	b := [...]string{"北京", "上海", "广州", "深圳"}

	for i := 0; i < len(b); i++ {
		fmt.Println("方式2，第", b[i], "次循环")
	}

	for k, v := range b {
		fmt.Println("方式3，第", k, "次循环", "值为", v)
	}
}
