package main

import (
	"fmt"
	"time"
)

func halo() {
	fmt.Println("Halo dari goroutine!")
}

func main() {
	go halo()
	time.Sleep(1 * time.Second)
	fmt.Println("Program selesai!")
}