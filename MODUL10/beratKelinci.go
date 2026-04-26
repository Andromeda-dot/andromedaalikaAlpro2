package main

import "fmt"

func main() {
	var beratKelinci [1000]float64
	var n, i int
	var beratTerkecil, beratTerbesar float64

	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil = beratKelinci[0]
	beratTerbesar = beratKelinci[0]

	for i = 1; i < n; i++ {
		if beratKelinci[i] < beratTerkecil {
			beratTerkecil = beratKelinci[i]
		}
		if beratKelinci[i] > beratTerbesar {
			beratTerbesar = beratKelinci[i]
		}
	}
	fmt.Println("Berat kelinci terkecil dan terbesar adalah:", beratTerkecil, beratTerbesar)
}
