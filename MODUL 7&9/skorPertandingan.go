package main

import "fmt"

func cariSkor(klubA, klubB string, skorA, skorB int) string {
	if skorA > skorB {
		return klubA
	} else if skorB > skorA {
		return klubB
	}
	return "Draw"
}

func main() {
	var klubA, klubB, hasil string
	var skorA, skorB, i, j, jumlahMenang int
	var pemenang [1000]string

	jumlahMenang = 0
	i = 1

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}
		hasil = cariSkor(klubA, klubB, skorA, skorB)

		pemenang[jumlahMenang] = hasil
		jumlahMenang++

		i++
	}

	for j = 0; j < jumlahMenang; j++ {
		fmt.Printf("Hasil %d: %s\n", j+1, pemenang[j])
	}
	fmt.Println("Pertandingan selesai")
}
