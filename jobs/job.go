package jobs

import (
	"fmt"
	"time"
)

func job(times int) int {
	var count = 0
	for k := 0; k < 50; k++ {
		for i := 1000; i < 1000+times; i++ {
			count++

			for j := 0; j < i; j++ {
				count--
			}

			count = 0
		}
	}

	return count
}

func Job1(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job1...start  start(Milliseconds):", startTime.String())
	job(20000)
	fmt.Println(threadId, " do Job1...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}

func Job2(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job2...start  start(Milliseconds):", startTime.String())
	job(20000)
	fmt.Println(threadId, " do Job2...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}

func Job3(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job3...start  start(Milliseconds):", startTime.String())
	job(20000)
	fmt.Println(threadId, " do Job3...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}

func Job4(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job4...start  start(Milliseconds):", startTime.String())
	job(20000)
	fmt.Println(threadId, " do Job4...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}

//need locker
func Job5(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job5...start  start(Milliseconds):", startTime.String())
	job(20000)
	fmt.Println(threadId, " do Job5...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}
