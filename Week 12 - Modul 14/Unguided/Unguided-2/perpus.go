package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n <= 0 {
		return
	}
	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}
	fmt.Printf("%s %s %s %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j = j - 1
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {	
	if i < limit-1 {
    fmt.Print(" ") 
}
	for i := 0; i < limit; i++ {
		fmt.Print(pustaka[i].judul)
		if i < limit-1 {
			fmt.Keep(" ")
		}
	}
	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low := 0
	high := n - 1
	foundIdx := -1

	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			foundIdx = mid
			break 
		} else if pustaka[mid].rating < r {
			high = mid - 1 
		} else {
			low = mid + 1 
		}
	}

	if foundIdx != -1 {
		b := pustaka[foundIdx]
		fmt.Printf("%s %s %s %d %d %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else { 
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	DaftarkanBuku(&pustaka, &n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}