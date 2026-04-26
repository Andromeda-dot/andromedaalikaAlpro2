package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	var i int
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i = 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var i int
	var total float64
	for i = 0; i < n; i++ {
		total += arrBerat[i]
	}
	if n == 0 {
		return 0
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n, i int
	var rata2, min, max float64

	fmt.Print("Masukkan data berat balita: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan berat balika ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)
	rata2 = rerata(data, n)

	fmt.Printf("Berat balita minimun: %.2f kg\n", min)
	fmt.Printf("Berat balita maximum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata2)
}
