package main

import "fmt"

func main() {
	fmt.Println("Loop 1: for i := 0; i < 5; i++")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("\nLoop 2: for i < 5 (while-style loop)")
	i := 0
	for i < 5 {
		fmt.Println("i =", i)
		i++
	}

	fmt.Println("\nLoop 3: infinite loop with break")
	counter := 0
	for {
		fmt.Println("counter =", counter)
		counter++
		if counter >= 3 {
			break
		}
	}

	fmt.Println("\nLoop 4: for-range over a slice")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", index, fruit)
	}
}
