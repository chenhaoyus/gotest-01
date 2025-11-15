package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Hello world")

	var a = 100

	fmt.Println("===========指针一==============")
	aa := test01(&a)

	fmt.Println(aa)

	fmt.Println("===========指针二==============")

	var b = []int{1, 2, 3, 4}
	bb := test02(&b)

	fmt.Println(bb)

	fmt.Println("===========go 协程 ==============")
	test03()

	fmt.Println("===========Goroutine ==============")
	test04()

	fmt.Println("===========channl1==============")
	test05()
	fmt.Println("===========channl2==============")
	test06()
	fmt.Println("===========锁机制1==============")
	test07()
	fmt.Println("===========锁机制2==============")
	test08()
}

func test01(count *int) int {

	var c = *count

	for i := 0; i < 10; i++ {
		c++
	}

	return c
}

func test02(b *[]int) []int {
	var c = *b

	for i := 0; i < len(c); i++ {
		c[i] += 2
	}

	return c
}

func test03() {

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()
		for i := 1; i < 10; i += 2 {
			fmt.Println("go-1", i)
		}
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for i := 0; i < 10; i += 2 {
			fmt.Println("go-2", i)
		}
	}()

	time.Sleep(time.Second)
	wg.Wait()
}

func test04() {

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		j := rand.Intn(10)
		go func(st int) {
			defer wg.Done()
			startTime := time.Now()
			time.Sleep(time.Duration(j) * time.Second)

			fmt.Println("go-"+strconv.Itoa(i), time.Since(startTime))

		}(j)

	}

	time.Sleep(time.Duration(2) * time.Second)
	wg.Wait()
}

func test05() {
	var c1 = make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(c2 chan<- int) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c2 <- i
		}

		close(c1)
		fmt.Println("chan 关闭")
	}(c1)

	wg.Add(1)
	go func(c2 <-chan int) {
		defer wg.Done()
	loop:
		for {
			select {
			case j, ok := <-c2:
				fmt.Println(j, ok)
				if !ok {
					fmt.Println("chan 已关闭")
					break loop
				}
			}
		}
	}(c1)

	time.Sleep(time.Duration(2) * time.Second)
	wg.Wait()

}

func test06() {
	var c1 = make(chan int, 100)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(c2 chan<- int) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			c2 <- i
		}

		close(c1)
		fmt.Println("chan 关闭")
	}(c1)

	wg.Add(1)
	go func(c2 <-chan int) {
		defer wg.Done()
	loop:
		for {
			select {
			case j, ok := <-c2:
				fmt.Println(j, ok)
				if !ok {
					fmt.Println("chan 已关闭")
					break loop
				}
			}
		}
	}(c1)

	time.Sleep(time.Duration(2) * time.Second)
	wg.Wait()

}

func test07() {

	var safeCount = &safeCount{0, sync.Mutex{}}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				safeCount.add()
			}
		}()
	}

	time.Sleep(time.Duration(2) * time.Second)
	wg.Wait()
	fmt.Println(safeCount.count)
}

type safeCount struct {
	count int
	mu    sync.Mutex
}

func (s *safeCount) add() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
}

func test08() {
	var count int32 = 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}

	time.Sleep(time.Duration(2) * time.Second)
	wg.Wait()
	fmt.Println(count)
}
