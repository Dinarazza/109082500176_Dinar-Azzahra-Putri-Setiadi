package main

import (
	"fmt"
)

func main() {
	var N int
	var berat [1000]float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println("Berat kelinci terkecil:", min)
	fmt.Println("Berat kelinci terbesar:", max)
}
