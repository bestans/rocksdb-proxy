package main

import (
	"bufio"
	"fmt"
	"os"
	"rdb/dll"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

func main() {
	//dll.Debug()
	dll.LoadDB()
	//dll.Test6()
	test1()
}
func test2()  {
	fmt.Println("val", dll.B3())
}
func test1()  {
	times, _ := strconv.Atoi(os.Args[1])
	//times := 1
	var value int64
	wait := sync.WaitGroup{}
	wait.Add(3)
	go func() {
		for j := 0; j < times; j++ {
			for i := 0; i < 100000; i++ {
				atomic.AddInt64(&value, int64(dll.B3()))
			}

			fmt.Println("times:", value)
			//time.Sleep(time.Millisecond * 10)
		}
		wait.Done()
	}()
	go func() {
		for j := 0; j < times; j++ {
			for i := 0; i < 100000; i++ {
				atomic.AddInt64(&value, int64(dll.B3()))
			}
			fmt.Println("times:", value)
			//time.Sleep(time.Millisecond * 10)
		}
		wait.Done()
	}()
	go func() {
		for j := 0; j < times; j++ {
			for i := 0; i < 100000; i++ {
				atomic.AddInt64(&value, int64(dll.B3()))
			}
			runtime.GC()
			fmt.Println("times:", value)
			//time.Sleep(time.Millisecond * 10)
		}
		wait.Done()
	}()
	wait.Wait()
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(value)
}