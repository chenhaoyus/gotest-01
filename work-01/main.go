package main

import (
	"fmt"
)

func main() {

	//只出现一次的数字
	c := test01([]int{11, 12, 12, 14, 15, 16, 17, 18, 19, 20, 11})
	fmt.Println(c)
	//回文数
	b := test02(10)

	fmt.Println(b)

}

func test01(a []int) []int {
	fmt.Println("===============只出现一次的数字==================")

	var b map[int]int = make(map[int]int)

	for i := range a {
		value, exists := b[a[i]]
		if exists {
			b[a[i]] = value + 1
		} else {
			b[a[i]] = 1
		}
	}

	c := []int{}

	for bkey, bvalue := range b {
		if bvalue > 1 {
			c = append(c, bkey)
		}
	}

	return c
}

func test02(x int) bool {
	fmt.Println("===============回文数==================")

	if x < 0 {
		return false
	}

	var b []int = make([]int, 0)

	for x > 0 {
		rem := x % 10
		x = x / 10
		b = append(b, rem)
	}

	for intdex, value := range b {
		if value != b[len(b)-intdex-1] {
			return false
		}
	}

	return true
}
