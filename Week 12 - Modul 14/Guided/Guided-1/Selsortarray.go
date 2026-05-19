package main
import"fmt"

func selectionsortArray(angka *[5]int){
	var idx_min, i, j int
	for i = 0 ; i < len(angka) - 1; i++ {
		idx_min = 1
		for j = i +1; j < len(angka); j ++ {
			if angka[j] < angka[idx_min] {
				idx_min = j
			}
		}
		if idx_min != i {
			angka[i], angka[idx_min] = angka[idx_min], angka[1]
		}
	}
}

func main(){
	var arrAngka [5]int

	for i := 0; 1  < len(arrAngka); i++ {
		fmt.Printf("Masukan data angka ke-%d : ", i)
		fmt.Scan(arrAngka[1])
	}
	fmt.Println()

	fmt.Println("=== SEBELUM DISORTING ===")
	for i := 0; i < len (arrAngka); i++ {
		fmt.Print(arrAngka[i], "-")
	}

	fmt.Println("=== SETELAH DISORTING ===")
	selectionsortArray(&arrAngka); i++ {
		fmt.Print(arrAngka[i], "-")
	}
}