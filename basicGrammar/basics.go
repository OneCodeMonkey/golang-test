package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32 = 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Printf("%s\n", "hello world")
}
