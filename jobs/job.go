package jobs

import (
	"fmt"
	"io"
	"os"
	"time"
)

//var file *os.File
//
//func init() {
//	file, _ = os.Create("/tmp/demo.txt")
//}

func copyFile1(srcFile, destFile string) error {
	file1, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file1.Close()
	defer file2.Close()
	//拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			fmt.Println("报错了。。。")
			return err
		}
		file2.Write(bs[:n])
	}
	return nil
}

func job(times int) int {
	arr := make([]byte, times)
	count := 1
	for k := 0; k < 4000; k++ {
		for i := 0; i < times/2; i++ {
			for j := 0; j < i; j++ {
				arr[j] = byte(j)
				count++
			}
		}
	}
	_ = copyFile1("/tmp/demo.txt", "/tmp/demo_out.txt")

	return count
}

const spendTime = 2000

func Job1(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job1...start  start(Milliseconds):", startTime.String())
	job(spendTime)
	fmt.Println(threadId, " do Job1...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}

func Job2(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job2...start  start(Milliseconds):", startTime.String())
	job(spendTime)
	fmt.Println(threadId, " do Job2...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}

func Job3(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job3...start  start(Milliseconds):", startTime.String())
	job(spendTime)
	fmt.Println(threadId, " do Job3...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}

func Job4(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job4...start  start(Milliseconds):", startTime.String())
	job(spendTime)
	fmt.Println(threadId, " do Job4...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}

//need locker
func Job5(threadId int) {
	startTime := time.Now()
	//fmt.Println("do Job5...start  start(Milliseconds):", startTime.String())
	job(spendTime)
	fmt.Println(threadId, " do Job5...end  spend(Milliseconds):", time.Since(startTime).Milliseconds())
}
