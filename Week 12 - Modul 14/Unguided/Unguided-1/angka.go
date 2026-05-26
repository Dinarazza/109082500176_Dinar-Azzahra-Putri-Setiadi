package main

import (
	"fmt"
)

func main() {
	var input int
	var data []int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	if len(data) == 0 {
		return
	}

	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}

	for i, val := range data {
		fmt.Printf("%d", val)
		if i < len(data)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	isConstant := true
	var diff int

	if len(data) > 1 {
		diff = data[1] - data[0]
		for i := 1; i < len(data)-1; i++ {
			if data[i+1]-data[i] != diff {
				isConstant = false
				break
			}
		}
	} else {
		diff = 0
	}

	if isConstant {
		fmt.Printf("Data berjarak %d\n", diff)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
