package main

import (
	"bufio"
	"fmt"
	"os"
	"rdb/dll"
	"sync"
	"sync/atomic"
)

func main() {
	test1()
}
func test2()  {
	fmt.Println("val", dll.B4())
}
func test1()  {
	//times, _ := strconv.Atoi(os.Args[1])
	times := 1
	var value int32
	wait := sync.WaitGroup{}
	wait.Add(3)
	go func() {
		for j := 0; j < times; j++ {
			for i := 0; i < 100000; i++ {
				atomic.AddInt32(&value, int32(dll.B3()))
			}

			fmt.Println("times:", value)
			//time.Sleep(time.Millisecond * 10)
		}
		wait.Done()
	}()
	go func() {
		for j := 0; j < times; j++ {
			for i := 0; i < 100000; i++ {
				atomic.AddInt32(&value, int32(dll.B3()))
			}
			fmt.Println("times:", value)
			//time.Sleep(time.Millisecond * 10)
		}
		wait.Done()
	}()
	go func() {
		for j := 0; j < times; j++ {
			for i := 0; i < 100000; i++ {
				atomic.AddInt32(&value, int32(dll.B3()))
			}
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