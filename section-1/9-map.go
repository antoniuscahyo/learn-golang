package main

import "fmt"

func main() {
	// 1. Membuat map kosong
	var scores map[string]int = make(map[string]int)

	// 2. Menambahkan data ke map
	scores["Alice"] = 90
	scores["Bob"] = 75
	scores["Charlie"] = 82

	// 3. Menampilkan map
	fmt.Println("Isi map scores:", scores)

	// 4. Mengakses nilai berdasarkan key
	fmt.Println("Nilai Bob:", scores["Bob"])

	// 5. Mengecek apakah key ada
	name := "David"
	value, exists := scores[name]
	if exists {
		fmt.Printf("Nilai %s: %d\n", name, value)
	} else {
		fmt.Printf("%s tidak ditemukan dalam map\n", name)
	}

	// 6. Looping map
	fmt.Println("\nLooping isi map:")
	for key, value := range scores {
		fmt.Printf("Nama: %s, Nilai: %d\n", key, value)
	}

	// 7. Menghapus item dari map
	delete(scores, "Charlie")
	fmt.Println("\nSetelah menghapus Charlie:", scores)
}
