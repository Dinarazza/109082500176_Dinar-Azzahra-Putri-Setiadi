package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.println(penjumlahan(n))
}

func penjumlahan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlhana(n-1)
	}
}
