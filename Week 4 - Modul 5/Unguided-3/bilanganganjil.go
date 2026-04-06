package main

import "fmt"

func ganjil(n int, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	ganjil(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Bilangan ganjil dari 1 sampai ", n, " adalah: ")
	ganjil(n, 1)
}
