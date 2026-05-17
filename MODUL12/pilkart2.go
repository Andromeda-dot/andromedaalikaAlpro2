package main

import "fmt"

const NMAX int = 21

type tabSuara [NMAX]int

func hitungSuara(calon *tabSuara, totalMasuk, suaraSah *int) {
	var x int
	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		*totalMasuk += 1

		if x >= 1 && x <= 20 {
			*suaraSah += 1
			calon[x] += 1
		}
	}
}

func cariPemenang(calon tabSuara, ketua, wakil *int) {
	var max1, max2, i int
	for i = 1; i <= 20; i++ {
		if calon[i] > max1 {
			max2 = max1
			*wakil = *ketua
			max1 = calon[i]
			*ketua = i
		} else if calon[i] == max1 && calon[i] > 0 {
			*wakil = i
		} else if calon[i] > max2 {
			max2 = calon[i]
			*wakil = i
		}
	}
}

func tampilkanHasil(totalMasuk, suaraSah, ketua, wakil int) {
	fmt.Println("Suara Masuk: ", totalMasuk)
	fmt.Println("Suara Sah: ", suaraSah)
	fmt.Println("Ketua RT: ", ketua)
	fmt.Println("Wakil RT: ", wakil)
}

func main() {
	var calon tabSuara
	var suaraSah, totalMasuk, ketua, wakil int

	hitungSuara(&calon, &totalMasuk, &suaraSah)
	cariPemenang(calon, &ketua, &wakil)
	tampilkanHasil(totalMasuk, suaraSah, ketua, wakil)
}
