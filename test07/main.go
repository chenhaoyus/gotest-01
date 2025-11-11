package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i := 0; i < len(a); i++ {
		fmt.Println("======", a[i], "==========")
	}

	for k, v := range a {
		fmt.Println("方式2，第", k, "次循环", "值为", v)
	}

	a = append(a, 110, 120, 130)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 110, 120, 130)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 110, 120, 130)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 110, 120, 130)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := make([]int, 5, 6)

	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = append(b, 20, 30, 40, 50, 60, 20, 30, 40, 50, 60)

	for k, v := range b {
		fmt.Println("方式3，第", k, "次循环", "值为", v)
	}

	fmt.Println(len(b))
	fmt.Println(cap(b))

}
