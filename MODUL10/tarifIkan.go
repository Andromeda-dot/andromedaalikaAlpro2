package main

import "fmt"

func main() {
	var x, y, i, jumlahIkan int
	var berat [1000]float64
	var totalWadah float64

	fmt.Scan(&x, &y)
	for i = 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	for i = 0; i < x; i++ {
		totalWadah += berat[i]
		jumlahIkan++
		if jumlahIkan == y || i == x-1 {
			fmt.Printf("%.2f ", totalWadah)
			totalWadah = 0
			jumlahIkan = 0
		}
	}
	fmt.Println()

	totalWadah = 0
	jumlahIkan = 0
	for i = 0; i < x; i++ {
		totalWadah += berat[i]
		jumlahIkan++
		if jumlahIkan == y || i == x-1 {
			fmt.Printf("%.2f ", totalWadah/float64(jumlahIkan))
			totalWadah = 0
			jumlahIkan = 0
		}
	}
	fmt.Println()
}
