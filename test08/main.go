package main

import "fmt"

func main() {
	fmt.Println("study map")

	var a = map[string]int{}

	a["张三"] = 18
	a["李四"] = 20
	a["王五"] = 25

	for k, v := range a {
		fmt.Println("key:", k, "value:", v)
	}

	b := map[int]string{1: "20", 2: "30", 3: "40"}

	for k, v := range b {
		fmt.Println("key:", k, "value:", v)
	}

	b[4] = "50"

	for k, v := range b {
		fmt.Println("key:", k, "value:", v)
	}

	delete(b, 3)

	for k, v := range b {
		fmt.Println("key:", k, "value:", v)
	}
}
