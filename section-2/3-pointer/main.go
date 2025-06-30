package main

import (
	"errors"
	"fmt"
)

func main() {
	tanpaPointer := 10
	var denganPointer *int = &tanpaPointer

	fmt.Println("Nilai x:", tanpaPointer)
	fmt.Println("Alamat x:", denganPointer)
	fmt.Println("Nilai dari pointer p:", *tanpaPointer)
	fmt.Println("=================")

	nilaiA := 200
	nilaiB := 300
	fmt.Println("Hasil Perhitungan:", hitungPenjumlahan(&nilaiA, &nilaiB))
	fmt.Println("=================")

	panjang := 5.0
	lebar := 10.0

	hasil, err := hitungLuasPersegiPanjang(panjang, lebar)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Panjang: %.2f, Lebar: %.2f\n", panjang, lebar)
	fmt.Printf("Luas Persegi Panjang:", *hasil)

}

func hitungPenjumlahan(a *int, b *int) int {
	return *a + *b
}

func hitungLuasPersegiPanjang(panjang, lebar float64) (*float64, error) {
	if panjang <= 0 || lebar <= 0 {
		return nil, errors.New("Panjang dan lebar harus lebih besar dari nol")
	}

	luas := panjang * lebar
	return &luas, nil
}