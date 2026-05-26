package main

import (
	"fmt"
)

func main() {
	var n, m, nomor int

	if fmt.Scan(&n); n <= 0 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			fmt.Scan(&nomor)
			if nomor%2 != 0 {
				ganjil = append(ganjil, nomor)
			} else {
				genap = append(genap, nomor)
			}
		}

		for step := 0; step < len(ganjil)-1; step++ {
			minIdx := step
			for k := step + 1; k < len(ganjil); k++ {
				if ganjil[k] < ganjil[minIdx] {
					minIdx = k
				}
			}
			ganjil[step], ganjil[minIdx] = ganjil[minIdx], ganjil[step]
		}

		for step := 0; step < len(genap)-1; step++ {
			maxIdx := step
			for k := step + 1; k < len(genap); k++ {
				if genap[k] > genap[maxIdx] {
					maxIdx = k
				}
			}
			genap[step], genap[maxIdx] = genap[maxIdx], genap[step]
		}

		var hasil []int
		hasil = append(hasil, ganjil...)
		hasil = append(hasil, genap...)

		for idx, val := range hasil {
			fmt.Printf("%d", val)
			if idx < len(hasil)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
