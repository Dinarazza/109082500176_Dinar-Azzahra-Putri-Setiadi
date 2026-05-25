package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_PEMAIN = 1001

type Pemain struct {
	Nama   string
	Gol    int
	Assist int
}

func selectionSort(data []Pemain, n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if data[j].Gol > data[idxMax].Gol {
				idxMax = j
			} else if data[j].Gol == data[idxMax].Gol {
				if data[j].Assist > data[idxMax].Assist {
					idxMax = j
				}
			}
		}
		if idxMax != i {
			data[i], data[idxMax] = data[idxMax], data[i]
		}
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	daftarPemain := make([]Pemain, n)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		var gol, assist int
		var nama string

		if scanner.Scan() {
			baris := scanner.Text()
			if baris == "" {
				scanner.Scan()
				baris = scanner.Text()
			}

			idxSpasiAssist := len(baris) - 1
			for idxSpasiAssist >= 0 && baris[idxSpasiAssist] != ' ' {
				idxSpasiAssist--
			}
			fmt.Sscanf(baris[idxSpasiAssist+1:], "%d", &assist)

			idxSpasiGol := idxSpasiAssist - 1
			for idxSpasiGol >= 0 && baris[idxSpasiGol] != ' ' {
				idxSpasiGol--
			}
			fmt.Sscanf(baris[idxSpasiGol+1:idxSpasiAssist], "%d", &gol)

			nama = baris[:idxSpasiGol]

			daftarPemain[i] = Pemain{Nama: nama, Gol: gol, Assist: assist}
		}
	}

	selectionSort(daftarPemain, n)

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", daftarPemain[i].Nama, daftarPemain[i].Gol, daftarPemain[i].Assist)
	}
}
