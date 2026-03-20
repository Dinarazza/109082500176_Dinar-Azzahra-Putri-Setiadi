package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c : ")
	fmt.Scan(&c)

	fmt.Printf("F(G(H( %d ))) : %d\n", a, f(g(h(a))))

	fmt.Printf("G(H(F( %d ))) : %d\n", b, g(h(f(b))))

	fmt.Printf("H(F(G( %d ))) : %d\n", c, h(f(g(c))))
}
