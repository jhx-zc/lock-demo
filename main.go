package main

import (
	"fmt"
	"lock_demo/demo1"
	"lock_demo/demo2"
	"sync"
	"time"
)

const (
	threadCount = 20
)

func testDemo() {
	startTime := time.Now()

	for i := 0; i < threadCount; i++ {
		demo1.DoWork(nil, i)
	}

	fmt.Println("demo1 spend time(Milliseconds):", time.Since(startTime).Milliseconds())
}

func testDemo1() {
	wg := &sync.WaitGroup{}

	startTime := time.Now()

	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go demo1.DoWork(wg, i)
	}

	wg.Wait()

	fmt.Println("demo1 spend time(Milliseconds):", time.Since(startTime).Milliseconds())
}

func testDemo2() {
	wg := &sync.WaitGroup{}

	startTime := time.Now()

	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go demo2.DoWork(wg, i)
	}

	wg.Wait()

	fmt.Println("demo2 spend time(Milliseconds):", time.Since(startTime).Milliseconds())
}

func main() {
	testDemo()

	fmt.Println("-----------")

	testDemo1()

	fmt.Println("-----------")

	testDemo2()
}
