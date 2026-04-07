package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
	return 0
}

func biayaPerHari(jumlahMhs int, tujuan string) int {
	// biaya domestik per mahasiswa per hari
	makan := 2 * 35000
	penginapan := 250000
	uangSaku := 300000

	biayaDomestik := makan + penginapan + uangSaku

	if tujuan == "domestik" {
		return biayaDomestik * jumlahMhs
	} else if tujuan == "mancanegara" {
		return int(float64(biayaDomestik) * 1.5 * float64(jumlahMhs))
	}
	return 0
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)
	biayaHarian := biayaPerHari(jumlahMhs, tujuan)

	*totalBiaya = float64(hariDitanggung * biayaHarian)
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Println(" Biaya perjalanan yang harus dikeluarkan Tel - U:", biaya)
}
