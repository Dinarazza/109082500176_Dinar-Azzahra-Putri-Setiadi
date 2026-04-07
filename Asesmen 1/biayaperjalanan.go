package main

import (
	"fmt"
	"strings"
)

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
	makan := 2 * 35000
	penginapan := 250000
	uangSaku := 300000

	biayaDomestik := makan + penginapan + uangSaku

	if tujuan == "domestik" {
		return biayaDomestik * jumlahMhs
	} else {
		return int(float64(biayaDomestik) * 1.5 * float64(jumlahMhs))
	}
}

func perhitunganBiaya(jumlahMhs, lama int, tujuan string, total *float64) {
	hari := tanggunganHari(lama, tujuan)
	biaya := biayaPerHari(jumlahMhs, tujuan)

	*total = float64(hari * biaya)
}

func formatRupiah(angka float64) string {
	return fmt.Sprintf("Rp%.0f", angka)
}

func main() {
	var jumlah, lama int
	var tujuan string
	var total float64

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)
	if jumlah <= 0 {
		fmt.Println("Jumlah mahasiswa harus lebih dari 0!")
		return
	}

	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)
	if lama <= 0 {
		fmt.Println("Lama perjalanan harus lebih dari 0!")
		return
	}

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)
	tujuan = strings.ToLower(tujuan)

	if tujuan != "domestik" && tujuan != "mancanegara" {
		fmt.Println("Tujuan harus 'domestik' atau 'mancanegara'!")
		return
	}

	perhitunganBiaya(jumlah, lama, tujuan, &total)

	fmt.Println("\n Biaya perjalanan yang harus dikeluarkan Tel - U")
	fmt.Println("Biaya perjalanan yang harus dikeluarkan Tel-U:", formatRupiah(total))
}
