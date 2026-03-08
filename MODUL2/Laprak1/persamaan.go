package main

import "fmt"

func main() {
	var K, pembilang, penyebut, fk, k, hasil float64
	fmt.Print("Nilai K: ")
	fmt.Scan(&K)

	hasil = 1.0

	for k = 0; k <= K; k++ {
		fk = float64(k)
		pembilang = ((4 * fk) + 2) * ((4 * fk) + 2)
		penyebut = ((4 * fk) + 1) * ((4 * fk) + 3)
		hasil = hasil * (pembilang / penyebut)
	}
	fmt.Printf("Nilai akar 2: %.10f\n", hasil)
}
