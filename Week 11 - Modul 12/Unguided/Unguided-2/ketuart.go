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

	var totalSuaraMasuk, totalSuaraSah int
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

	ketua, wakil := -1, -1
	maks1, maks2 := -1, -1

	for i := 1; i <= 20; i++ {
		suara := perolehanSuara[i]
		if suara > maks1 {
			maks2 = maks1
			wakil = ketua
			maks1 = suara
			ketua = i
		} else if suara > maks2 {
			maks2 = suara
			wakil = i
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)
	if ketua != -1 {
		fmt.Printf("Ketua RT: %d\n", ketua)
	}
	if wakil != -1 {
		fmt.Printf("Wakil ketua: %d\n", wakil)
	}
}
