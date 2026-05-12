package main

import "fmt"

type arrData [5]string

func seqSearch(arr arrData, namaBintang string) int {
	var found bool = false
	for i := 0; i < len(arr); i++ {
		if arr[i] == Bintangcari {
			found = i
			break
		}
	}
	return found
}

func main() {
	var arrBintang arrData

	for i := 0; i < len(arrBintang); i++ {
		fmt.Printf("Masukan data bintang ke-%d:", i)
		fmt.Scan(&arrBintang[i])
	}
	fmt.Println()

	var bintangCari string
	fmt.Print("Masukan nama bintang yang mau dicari :")
	fmt.Scan(&bintangCari)

	var idxCari int
	idxCari = seqSearch(arrBintang, bintangCari)

	if idxCari > -1 {
		fmt.Printf("Data %s ditemukan pada indeks ke-%d!", bintangCari, idxCari)
	} else if idxCari == -1 {
		fmt.Printf("Data %s tidak ditemukan!", bintangCari)
	}
}
