package main

import "fmt"

type arrData [5]string

func seqSearch(arr arrData, namaBinatang string) int {
	var found int = false
	for i := 0; i < len(arr); i++ {
		if arr[i] == binatangcari {
			found = i
			break
		}
	}
	return found
}

func main() {
	var arrBintang arrData

	for i := 0; i < len(arrBinatang); i++ {
		fmt.Printf("Masukan data bintang ke-%d:", i)
		fmt.Scan(&arrBinatang[i])
	}
	fmt.Println()

	var binatangCari string
	fmt.Print("Masukan nama binatang yang mau dicari :")
	fmt.Scan(&bintangCari)

	var idxCari int
	idxCari = seqSearch(arrBinatang, binatangCari)

	if idxCari > -1 {
		fmt.Printf("Data %s ditemukan pada indeks ke-%d!", binatangCari, idxCari)
	} else if idxCari == -1 {
		fmt.Printf("Data %s tidak ditemukan!", binatangCari)
	}
}
