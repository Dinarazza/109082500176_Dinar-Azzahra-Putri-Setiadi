package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

func inputSet(T *set, n *int) {
	*n = 0
	var val int
	for {
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}
		if exist(*T, *n, val) {
			break
		}
		(*T)[*n] = val
		*n++
		if *n >= len(T) {
			break
		}
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		val := T1[i]
		if exist(T2, m, val) {
			(*T3)[*h] = val
			(*h)++
		}
	}
}

func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d", T[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}
