package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	jobs := make(chan int, 10)
	for {
		jobs <- 1
		go worker(jobs)
	}
}

func worker(jobs chan int) {
	fmt.Printf("Number of goroutine: %d, Number of CPU: %d\n", runtime.NumGoroutine(), runtime.NumCPU())
	time.Sleep(1 * time.Second)
	<-jobs
}
