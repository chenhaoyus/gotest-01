package main

import (
	"fmt"
)

func main() {
	fmt.Println("test04")
	aaaa(A1)
}

type a string

const (
	A1 a = "A1"
	A2 a = "A2"
	A3 a = "A3"
)

func aaaa(a1 a) {
	fmt.Println(a1)
}
