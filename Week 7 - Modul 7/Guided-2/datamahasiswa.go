package main

import "fmt"

type kalimat string

type mahasiswa struct {
	nama  kalimat
	nim   int
	nilai float64
}

func main() {
	var m mahasiswa

	fmt.Print("Masukan nama mahasiswa: ")
	fmt.Scan(&m.nama)
	fmt.Print("Masukan nim: ")
	fmt.Scan(&m.nim)
	fmt.Print("Masukan nilai: ")
	fmt.Scan(&m.nilai)

	fmt.Println("DATA MAHASISWA")
	fmt.Println("nama:", m.nama)
	fmt.Println("nim:", m.nim)
	fmt.Println("nilai:", m.nilai)
}
