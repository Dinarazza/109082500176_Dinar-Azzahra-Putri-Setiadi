package main

import (
	"fmt"
)

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := x / y
	var totalPerWadah []float64
	var totalSemua float64

	index := 0

	for i := 0; i < jumlahWadah; i++ {
		var total float64
		for j := 0; j < y; j++ {
			total += ikan[index]
			index++
		}
		totalPerWadah = append(totalPerWadah, total)
		totalSemua += total
	}

	for i := 0; i < len(totalPerWadah); i++ {
		fmt.Printf("%.2f ", totalPerWadah[i])
	}
	fmt.Println()

	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rataRata)
}
