package main

import (
	"fmt"
	"math"
)

const MAX int = 100

func isiArray(arr [MAX]int, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func rata2(arr [MAX]int, n int) float64 {
	var total, i int
	for i = 0; i < n; i++ {
		total += arr[i]
	}
	return float64(total) / float64(n)
}

func deviasi(arr [MAX]int, n int, rata float64) float64 {
	var jumlahKuadrat float64
	var i int
	for i = 0; i < n; i++ {
		jumlahKuadrat += math.Pow(float64(arr[i])-rata, 2)
	}
	return math.Sqrt(jumlahKuadrat / float64(n))
}

func hitungFrekuensi(arr [MAX]int, n int, target int) int {
	var hitung, i int
	hitung = 0
	for i = 0; i < n; i++ {
		if arr[i] == target {
			hitung++
		}
	}
	return hitung
}

func hapusElemen(arr *[MAX]int, n int, elemen int) int {
	var i int
	for i = elemen; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	return n - 1
}

func main() {
	var arr [MAX]int
	var n, x, elemenHapus, targetFrekuensi, i int
	var rata, sd float64

	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	fmt.Printf("Masukkan %d angka: ", n)
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("a. Isi keseluruhan array: ")
	isiArray(arr, n)

	fmt.Print("b. Indeks ganjil: ")
	for i = 1; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("c. Indeks genap: ")
	for i = 0; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("d. Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen indeks kelipatan %d: ", x)
	for i = 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	rata = rata2(arr, n)
	fmt.Printf("f. Rata-rata: %.2f\n", rata)

	sd = deviasi(arr, n, rata)
	fmt.Printf("g. Standar deviasi: %.2f\n", sd)

	fmt.Print("h. Cari frekuensi angka: ")
	fmt.Scan(&targetFrekuensi)
	fmt.Printf("Angka %d muncul %d kali\n", targetFrekuensi, hitungFrekuensi(arr, n, targetFrekuensi))

	fmt.Print("e. Indeks yang ingin dihapus: ")
	fmt.Scan(&elemenHapus)
	n = hapusElemen(&arr, n, elemenHapus)
	fmt.Print("Isi array sekarang: ")
	isiArray(arr, n)
}
