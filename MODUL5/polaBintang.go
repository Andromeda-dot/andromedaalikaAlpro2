package main

import "fmt"

func cetakBintang(jumlah int) {
	if jumlah <= 0 {
		return
	}
	fmt.Print("*")
	cetakBintang(jumlah - 1)
}

func pola(n int) {
	if n <= 0 {
		return
	}
	pola(n - 1)
	cetakBintang(n)
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	pola(n)
}
