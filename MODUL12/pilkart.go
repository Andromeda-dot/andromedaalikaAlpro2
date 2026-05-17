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

func tampilkanHasil(calon tabSuara, totalMasuk, suaraSah int) {
	var i int
	fmt.Println("Suara Masuk: ", totalMasuk)
	fmt.Println("Suara Sah: ", suaraSah)

	for i = 1; i <= 20; i++ {
		if calon[i] > 0 {
			fmt.Printf("%d: %d\n", i, calon[i])
		}
	}
}

func main() {
	var calon tabSuara
	var suaraSah, totalMasuk int

	hitungSuara(&calon, &totalMasuk, &suaraSah)
	tampilkanHasil(calon, totalMasuk, suaraSah)
}
