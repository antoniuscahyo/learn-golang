package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Major string
}

func main() {
	var student1 Student
	student1.Name = "Budi"
	student1.Age = 20
	student1.Major = "Teknik Informatika"

	student1 := Student{
		Name:  "siti",
		Age:   22,
		Major: "Sistem Informatika",
	}

	student3 := Student{
		Name:  "Andi",
		Age:   23,
		Major: "Teknik Komputer",
	}

	fmt.Println("Mahasiswa 1:", student1)
	fmt.Println("Mahasiswa 2:", student1)
	fmt.Println("Mahasiswa 3:", student3)

	fmt.Println("Nama Mahasiswa 1:", student1.Name)

	student1.Age = 21
	fmt.Println("Umur Mahasiswa 1:", student1.Age)

	printStudent(student2)
}

func printStudent(student Student) {
	fmt.Printf("Nama: %s, Umur: %d, Jurusan: %s\n", student.Name, student.Age, student.Major)
}