package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string
	pertandingan := 1

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		// berhenti jika ada skor negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Println("Hasil:", klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			fmt.Println("Hasil:", klubB)
		} else {
			pemenang = append(pemenang, "Draw")
			fmt.Println("Hasil: Draw")
		}

		pertandingan++
	}

	fmt.Println("\nDaftar hasil pertandingan:")
	for i, h := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, h)
	}

	fmt.Println("Pertandingan selesai")
}
