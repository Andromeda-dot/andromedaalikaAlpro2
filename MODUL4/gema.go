package main

import "fmt"

func hitungSkor(soal, skor *int) {
	var i, waktu int
	*soal = 0
	*skor = 0

	for i = 1; i <= 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soalNow, skorNow, maxSoal, minSkor int
	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&soalNow, &skorNow)

		if soalNow > maxSoal {
			maxSoal = soalNow
			minSkor = skorNow
			pemenang = nama
		} else if soalNow == maxSoal {
			if skorNow < minSkor {
				minSkor = skorNow
				pemenang = nama
			}
		}
	}

	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minSkor)
	}
}
