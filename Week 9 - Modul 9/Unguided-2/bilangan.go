package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, hapus, cari int
	var jumlah, rata, variansi, standarDeviasi float64

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	data := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan elemen ke-", i, ": ")
		fmt.Scan(&data[i])
	}

	fmt.Print("Masukkan indeks kelipatan x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapus)

	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)

	fmt.Println("\nIsi keseluruhan array:")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}

	fmt.Println("\n\nElemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}

	fmt.Println("\n\nElemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}

	fmt.Println("\n\nElemen dengan indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(data[i], " ")
		}
	}

	fmt.Println("\n\nArray setelah indeks", hapus, "dihapus:")
	for i := hapus; i < n-1; i++ {
		data[i] = data[i+1]
	}
	n--
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}

	for i := 0; i < n; i++ {
		jumlah += float64(data[i])
	}
	rata = jumlah / float64(n)

	for i := 0; i < n; i++ {
		variansi += math.Pow(float64(data[i])-rata, 2)
	}
	variansi = variansi / float64(n)
	standarDeviasi = math.Sqrt(variansi)

	frekuensi := 0
	for i := 0; i < n; i++ {
		if data[i] == cari {
			frekuensi++
		}
	}

	fmt.Println("\n\nRata-rata:", rata)
	fmt.Println("Standar deviasi:", standarDeviasi)
	fmt.Println("Frekuensi bilangan", cari, ":", frekuensi)
}
