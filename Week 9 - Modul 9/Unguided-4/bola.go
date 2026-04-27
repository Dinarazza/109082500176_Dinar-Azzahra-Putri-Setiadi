package main

import "fmt"

const NMAX = 127

type Tabel struct {
	tab [NMAX]rune
	n   int
}

func isiArray(t *Tabel, n *int) {
	fmt.Scan(n)

	for i := 0; i < *n && i < NMAX; i++ {
		fmt.Scanf(" %c", &t.tab[i])
	}

	t.n = *n
}

func cetakArray(t Tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t.tab[i])
	}
	fmt.Println()
}

func balikArray(t *Tabel, n int) {
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-1-i] = t.tab[n-1-i], t.tab[i]
	}
}

func palindrom(t Tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t.tab[i] != t.tab[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab Tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Array awal: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}

	balikArray(&tab, n)

	fmt.Print("Array setelah dibalik: ")
	cetakArray(tab, n)
}
