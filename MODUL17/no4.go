package main

import "fmt"

func hitungSuku(i int) float64 {
	var pembilang float64
	var penyebut float64

	if i%2 != 0 {
		pembilang = 1.0
	} else {
		pembilang = -1.0
	}

	penyebut = float64(2*i - 1)

	return pembilang / penyebut
}

func hitungPiDeret(n int) {
	var totalDeret float64 = 0.0
	var hasilPi float64
	var i int

	for i = 1; i <= n; i++ {
		totalDeret = totalDeret + hitungSuku(i)
	}

	hasilPi = 4.0 * totalDeret

	fmt.Printf("Hasil Pi: %.7f\n", hasilPi)
}

func main() {
	var n int

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	hitungPiDeret(n)
}
