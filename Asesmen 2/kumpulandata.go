package main

import (
	"fmt"
)

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		var nama string
		var p int
		var t float64
		_, err := fmt.Scan(&nama, &p, &t)
		if err != nil {
			break
		}
		prov[i] = nama
		pop[i] = p
		tumbuh[i] = t
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := float64(pop[i]) * (tumbuh[i] + 1)
			fmt.Printf("%s %.0f\n", prov[i], prediksi)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	InputData(&prov, &pop, &tumbuh)

	var target string
	fmt.Scan(&target)

	idxTercepat := ProvinsiTercepat(tumbuh)
	idxDicari := IndeksProvinsi(prov, target)

	fmt.Println()
	fmt.Printf("Provinsi dengan angka pertumbuhan tercepat : %s\n", prov[idxTercepat])
	fmt.Println()
	if idxDicari != -1 {
		fmt.Printf("Data provinsi yang dicari : %s\n", prov[idxDicari])
	} else {
		fmt.Printf("Data provinsi yang dicari : %s (tidak ditemukan)\n", target)
	}
	fmt.Println()
	Prediksi(prov, pop, tumbuh)
}
