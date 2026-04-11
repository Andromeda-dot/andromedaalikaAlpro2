package main

import "fmt"

func bilanganGanjil(n int) {
	var i int
	for i = 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	var angka int
	fmt.Scan(&angka)
	bilanganGanjil(angka)
}
