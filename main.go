package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int, 10)
	for {
		ch <- 1
		go helper(ch)
	}
}

func helper(ch chan int) {
	fmt.Printf("Number of goroutine: %d, Number of CPU: %d\n", runtime.NumGoroutine(), runtime.NumCPU())
	time.Sleep(1 * time.Second)
	<-ch
}
