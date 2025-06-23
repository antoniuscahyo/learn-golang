package main

import "fmt"

func main() {
	// 1. Deklarasi array dengan panjang tetap
	var numbers [5]int

	// 2. Mengisi nilai array secara manual
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// 3. Menampilkan seluruh isi array
	fmt.Println("Isi array numbers:", numbers)

	// 4. Mengakses elemen array
	fmt.Println("Elemen pertama:", numbers[0])
	fmt.Println("Elemen terakhir:", numbers[len(numbers)-1])

	// 5. Looping menggunakan indeks
	fmt.Println("\nLooping array dengan indeks:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	// 6. Looping menggunakan for-range
	fmt.Println("\nLooping array dengan for-range:")
	for index, value := range numbers {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// 7. Deklarasi dan inisialisasi langsung
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("\nArray names:", names)
}
