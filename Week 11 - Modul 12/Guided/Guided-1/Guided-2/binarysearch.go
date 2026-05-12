package main

import "fmt"

func binarySearch(data []int, x int) int {

	kiri := 0

	kanan := len(data) - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if data[tengah] == x {
			return tengah

		} else if data[tengah] < x {
			kiri = tengah + 1

		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func main() {

	var n int
	var cari int

	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	data := make([]int, n)

	fmt.Println("Masukkan data terurut:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Print("Angka yang dicari: ")
	fmt.Scan(&cari)

	hasil := binarySearch(data, cari)

	if hasil != -1 {
		fmt.Println("Data ditemukan di index:", hasil)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
