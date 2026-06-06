package main

import "fmt"

func analisisString(x string, n int) (bool, int, int, bool) {
	var ketemu, minimalDua bool
	var posisiPertama, jumlahX, i int
	var currentString string

	ketemu = false
	minimalDua = false
	posisiPertama = -1
	jumlahX = 0

	fmt.Printf("Masukkan %d data string:\n", n)

	for i = 1; i <= n; i++ {
		fmt.Printf("Data ke-%d: ", i)
		fmt.Scan(&currentString)

		if currentString == x {
			ketemu = true
			jumlahX = jumlahX + 1

			if posisiPertama == -1 {
				posisiPertama = i
			}
		}
	}
	if jumlahX >= 2 {
		minimalDua = true
	}
	return ketemu, posisiPertama, jumlahX, minimalDua
}

func main() {
	var x string
	var n, posisi, totalX int
	var adaX, minimalDua bool

	fmt.Print("Masukkan X: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	adaX, posisi, totalX, minimalDua = analisisString(x, n)

	fmt.Println("HASIL ANALISIS:")
	if adaX == true {
		fmt.Println("a. Apakah string x ada dalam kumpulan n data string tersebut? Ya, ada")
	} else {
		fmt.Println("a. Apakah string x ada dalam kumpulan n data string tersebut? Tidak ada")
	}

	if posisi != 1 {
		fmt.Printf("b. Pada posisi ke berapa string x tersebut ditemukan? Posisi ke-%d\n", posisi)
	} else {
		fmt.Println("b. Pada posisi ke berapa string x tersebut ditemukan? Tidak ditemukan")
	}

	fmt.Printf("c. Ada berapakah string x dalam kumpulan n data string tersebut? Ada %d\n", totalX)

	if minimalDua == true {
		fmt.Println("d. Adakah sedikitnya dua string x dalam n data string tersebut? Ya, ada sedikitnya dua")
	} else {
		fmt.Println("d. Adakah sedikitnya dua string x dalam n data string tersebut? Tidak, kurang dari dua")
	}
}
