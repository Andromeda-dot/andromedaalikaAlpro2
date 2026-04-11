package main

import (
	"fmt"
)

func faktorBilangan(n int) {
	var i int
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	var angka int
	fmt.Scan(&angka)
	faktorBilangan(angka)
}
