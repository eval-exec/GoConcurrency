package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	in := make(chan func())

	wg.Add(1)
	go func(in chan func()) {
		for {
			fmt.Println("等待任务处理")
			f := <-in
			f()
		}
		wg.Done()
	}(in)

	for i := 0; i < 100; i++ {
		in <- func(i int) func() {
			return func() {
				fmt.Println(i)
			}
		}(i)
	}

	wg.Wait()
	// time.Sleep(10 * time.Second)
}
