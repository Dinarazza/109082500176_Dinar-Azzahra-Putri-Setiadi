package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()

	tokens := strings.Fields(line)

	var totalSuaraMasuk int
	var totalSuaraSah int

	var perolehanSuara [21]int

	for _, token := range tokens {
		angka, err := strconv.Atoi(token)
		if err != nil {
			continue
		}
		if angka == 0 {
			break
		}

		totalSuaraMasuk++

		if angka >= 1 && angka <= 20 {
			totalSuaraSah++
			perolehanSuara[angka]++
		}
	}
	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)

	for i := 1; i <= 20; i++ {
		if perolehanSuara[i] > 0 {
			fmt.Printf("%d: %d\n", i, perolehanSuara[i])
		}
	}
}
