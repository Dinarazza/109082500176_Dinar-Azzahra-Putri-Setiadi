package main

import "fmt"

func main() {
	var nilai map[string]int = make(map[string]int)

	nilai["Dhimas"] = 90
	nilai["Ichya"] = 90

	fmt.Println("Data nilai : ")
	var nama string
	var grade int
	for nama, grade = range nilai {
		fmt.Println(nama, " = ", grade)
	}

	nilai["Ichya"] = 80

	//operasi searching
	var cariNama string = "hafizh"
	var isiData int
	var ok bool
	isiData, ok = nilai[cariNama]
	if ok {
		fmt.Println("Nilai", cariNama, " = ", isiData)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	delete(nilai, "Dhimas")
}
