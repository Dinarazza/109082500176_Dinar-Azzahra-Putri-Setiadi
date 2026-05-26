package main

import (
	"fmt"
)

func main() {
	var n, m int
	if fmt.Scan(&n); n <= 0 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		for step := 0; step < m-1; step++ {
			minIdx := step
			for k := step + 1; k < m; k++ {
				if rumah[k] < rumah[minIdx] {
					minIdx = k
				}
			}

			rumah[step], rumah[minIdx] = rumah[minIdx], rumah[step]
		}

		for idx, val := range rumah {
			fmt.Printf("%d", val)
			if idx < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
