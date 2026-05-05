package main

import "fmt"

const nMax = 51

type Mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}

type ArrayMahasiswa [nMax]Mahasiswa

func getFirstValue(arr ArrayMahasiswa, n int, targetNIM string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == targetNIM {
			return arr[i].Nilai
		}
	}
	return -1
}

func getMaxValue(arr ArrayMahasiswa, n int, targetNIM string) int {
	maxVal := -1
	found := false
	for i := 0; i < n; i++ {
		if arr[i].NIM == targetNIM {
			if !found || arr[i].Nilai > maxVal {
				maxVal = arr[i].Nilai
				found = true
			}
		}
	}
	return maxVal
}

func main() {
	var arr ArrayMahasiswa
	var N int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&arr[i].NIM, &arr[i].Nama, &arr[i].Nilai)
	}

	var targetNIM string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&targetNIM)

	firstVal := getFirstValue(arr, N, targetNIM)
	maxVal := getMaxValue(arr, N, targetNIM)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", targetNIM, firstVal)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", targetNIM, maxVal)
}
