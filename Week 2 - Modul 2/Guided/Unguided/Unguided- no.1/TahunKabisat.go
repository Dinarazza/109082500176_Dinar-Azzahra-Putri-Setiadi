package main

import (
	"fmt"
)

func isKabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%4 == 0 && tahun%100 != 0 {
		return true
	}
	return false
}

func main() {
	var tahun int

	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&tahun)

	hasil := isKabisat(tahun)

	fmt.Println("Kabisat :", hasil)
}