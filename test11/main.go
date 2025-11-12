package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	ch := make(chan int, 6)

	a(ch)
	b(ch)
}

func a (ch chan<- int) {
	
	for i :=0;i<5;i++ {	
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}

	close(ch)
}

func b( ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
