package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func taskA(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task A started")
	time.Sleep(200 * time.Second)
	fmt.Println("Task A completed")
}

func taskB(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task B started")
	time.Sleep(3 * time.Second)
	fmt.Println("Task B completed")
}

func main() {
	runtime.GOMAXPROCS(4)

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)

	go taskA(&wg)
	go taskB(&wg)

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("All tasks completed in %v\n", duration)
	fmt.Println("Program finished!")	
}