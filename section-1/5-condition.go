package main

import "fmt"

func main() {
	// Contoh penggunaan if-else
	score := 75

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D or below")
	}

	// Contoh penggunaan if dengan short statement
	if x := 10; x%2 == 0 {
		fmt.Println("x adalah bilangan genap")
	} else {
		fmt.Println("x adalah bilangan ganjil")
	}

	// Contoh penggunaan switch
	day := "Saturday"

	switch day {
	case "Monday":
		fmt.Println("Hari Senin, semangat kerja!")
	case "Saturday", "Sunday":
		fmt.Println("Akhir pekan, waktunya liburan!")
	default:
		fmt.Println("Hari kerja biasa")
	}

	// Contoh switch tanpa ekspresi (seperti if-else)
	age := 17

	switch {
	case age < 13:
		fmt.Println("Anak-anak")
	case age < 18:
		fmt.Println("Remaja")
	default:
		fmt.Println("Dewasa")
	}

	// Simulasi ternary menggunakan if
	var status string
	if age >= 18 {
		status = "Dewasa"
	} else {
		status = "Belum dewasa"
	}
	fmt.Println("Status umur:", status)
}
