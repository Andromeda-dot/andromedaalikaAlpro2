package main

import "fmt"

func main() {
	var a, b float64
	for a < 9 && b < 9 {
		fmt.Print("Masukkan berat belanjaan di kedua kantong:")
		fmt.Scan(&a, &b)
	}
	fmt.Print("Proses selesai")
}
