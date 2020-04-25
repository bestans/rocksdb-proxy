package test

import (
	"sync"
	"testing"
)

var totalselecttimes = 10000000
func TestChannelSelect13Channel(t *testing.T) {
	ch := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	ch3 := make(chan bool, 1)
	ch4 := make(chan bool, 1)
	ch5 := make(chan bool, 1)
	ch6 := make(chan bool, 1)
	ch7 := make(chan bool, 1)
	ch8 := make(chan bool, 1)
	ch9 := make(chan bool, 1)
	ch10 := make(chan bool, 1)
	ch11 := make(chan bool, 1)
	ch12 := make(chan bool, 1)
	ch13 := make(chan bool, 1)
	wait := &sync.WaitGroup{}
	wait.Add(1)
	go func() {
		value := 0
		for {
			select {
			case <-ch:
				value++
				if value >= totalselecttimes {
					wait.Done()
					return
				}
			case <-ch2:
			case <-ch3:
			case <-ch4:
			case <-ch5:
			case <-ch6:
			case <-ch7:
			case <-ch8:
			case <-ch9:
			case <-ch10:
			case <-ch11:
			case <-ch12:
			case <-ch13:
			}
		}
	}()
	go func() {
		for i:= 0; i < totalselecttimes; i++ {
			ch <- true
		}
	}()
	wait.Wait()
}
func TestChannelSelect3Channel(t *testing.T) {
	ch := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	ch3 := make(chan bool, 1)
	wait := &sync.WaitGroup{}
	wait.Add(1)
	go func() {
		value := 0
		for {
			select {
			case <-ch:
				value++
				if value >= totalselecttimes {
					wait.Done()
					return
				}
			case <-ch2:
			case <-ch3:
			}
		}
	}()
	go func() {
		for i:= 0; i < totalselecttimes; i++ {
			ch <- true
		}
	}()
	wait.Wait()
}
func TestChannelSelect2Channel(t *testing.T) {
	ch := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	wait := &sync.WaitGroup{}
	wait.Add(1)
	go func() {
		value := 0
		for {
			select {
			case <-ch:
				value++
				if value >= totalselecttimes {
					wait.Done()
					return
				}
			case <-ch2:
			}
		}
	}()
	go func() {
		for i:= 0; i < totalselecttimes; i++ {
			ch <- true
		}
	}()
	wait.Wait()
}
func TestChannelSelect1Channel(t *testing.T) {
	ch := make(chan bool, 1)
	wait := &sync.WaitGroup{}
	wait.Add(1)
	go func() {
		value := 0
		for {
			select {
			case <-ch:
				value++
				if value >= totalselecttimes {
					wait.Done()
					return
				}
			}
		}
	}()
	go func() {
		for i:= 0; i < totalselecttimes; i++ {
			ch <- true
		}
	}()
	wait.Wait()
}

func TestMap(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	for id, _ := range m {
		if id % 2 == 0 {
			delete(m, id)
		}
	}
	for id, value := range m {
		println(id, value)
	}
}