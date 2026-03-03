package main
import "fmt"

func main () {
	var target, tabungan, total, hari int

	fmt.Print("Masukan target uang yang ingin dicapai: ")
	fmt.Scan(&target)

	total = 0
	hari = 0 

	for total < target {
		hari++
		fmt.Printf("Masukan nominal tabungan hari ke-%d: ", hari )
		fmt.Scan(&tabungan)

		total = total + tabungan
	}

	fmt.Printf("Selamat target tercapai dalam %d hari.\n", hari)
	fmt.Printf("Total tabungan Anda terkumpul: Rp%d\n", total)
}