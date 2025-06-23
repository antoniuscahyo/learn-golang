package main

import "fmt"

func main() {
	// 1. Membuat slice dari array
	arr := [5]int{10, 20, 30, 40, 50}
	slice1 := arr[1:4] // mengambil elemen index 1, 2, dan 3 (tidak termasuk 4)
	fmt.Println("Slice dari array:", slice1)

	// 2. Membuat slice langsung
	slice2 := []string{"apel", "jeruk", "mangga"}
	fmt.Println("Slice buah:", slice2)

	// 3. Menambahkan elemen ke slice (menggunakan append)
	slice2 = append(slice2, "pisang", "nanas")
	fmt.Println("Setelah append:", slice2)

	// 4. Mengambil panjang dan kapasitas slice
	fmt.Println("Panjang slice2:", len(slice2))
	fmt.Println("Kapasitas slice2:", cap(slice2))

	// 5. Looping slice dengan for-range
	fmt.Println("\nLooping slice buah:")
	for i, buah := range slice2 {
		fmt.Printf("Index %d: %s\n", i, buah)
	}

	// 6. Meng-copy slice
	source := []int{1, 2, 3}
	dest := make([]int, len(source))
	copy(dest, source)
	fmt.Println("\nHasil copy slice:", dest)
}
