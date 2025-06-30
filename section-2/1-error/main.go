package main

import (
	"fmt"
	"errors"
)

func validasiUmur(umur int) error {
	if umur < 0 {
		return errors.New("UMUR TIDAK BOLEH NEGATIF")
	} 
	if umur > 18 {
		return fmt.Errorf("UMUR %d BELUM CUKUP UNTUK MENDAFTAR", umur)
	}
	return nil
}

func main() {
	err := validasiUmur(15)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Pendaftaran berhasil!")
}