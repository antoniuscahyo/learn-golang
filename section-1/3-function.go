package main

import "fmt"

func calc(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff
}

func calcV2(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	s, d := calc(7, 5)
	fmt.Println("Jumlah:", s)
	fmt.Println("Selisih:", d)
}
