package main

import "fmt"

type desimal float64

func main() {
	var alas, tinggi, luas desimal

	fmt.Print("Masukan Pnajang Alas Segitiga: ")
	fmt.Scan(&alas)

	fmt.Print("Masukan Tinggi Segitiga: ")
	fmt.Scan(&tinggi)

	luas = 0.5 * alas * tinggi

	fmt.Printf("Luas Segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", alas, tinggi, luas)
}
