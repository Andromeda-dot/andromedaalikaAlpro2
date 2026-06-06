package main

import (
	"fmt"
	"math"
)

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

func simulasiKonvergensi() {
	var totalDeret float64 = 0.0
	var selisih, sukuSekarang, sukuBerikutnya, piSekarang, piBerikutnya float64
	var i int = 1

	for {
		sukuSekarang = hitungSuku(i)
		sukuBerikutnya = hitungSuku(i + 1)

		totalDeret = totalDeret + sukuSekarang

		selisih = math.Abs(sukuBerikutnya - sukuSekarang)

		if selisih <= 0.00001 && i > 1 {
			break
		}
		i += 1
	}
	fmt.Printf("Hasil PI: %.10f\n", piSekarang)
	fmt.Printf("Hasil PI: %.10f\n", piBerikutnya)
	fmt.Printf("Pada i ke: %d\n", i)
}

func main() {
	var n int

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	simulasiKonvergensi()
}
