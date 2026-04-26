package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func IPK_2(T arrMhs, n int) int {
	var idx int = 0
	var j int = 1
	for j < n {
		if T[idx].ipk < T[j].ipk {
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var T arrMhs
	var n, i, idxTerbaik int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("\nData Mahasiswa ke-%d\n", i+1)
		fmt.Print("Nama    : ")
		fmt.Scan(&T[i].nama)
		fmt.Print("NIM     : ")
		fmt.Scan(&T[i].nim)
		fmt.Print("Kelas   : ")
		fmt.Scan(&T[i].kelas)
		fmt.Print("Jurusan : ")
		fmt.Scan(&T[i].jurusan)
		fmt.Print("IPK     : ")
		fmt.Scan(&T[i].ipk)
	}

	if n > 0 {
		idxTerbaik = IPK_2(T, n)

		fmt.Println("\n--- Mahasiswa dengan IPK Tertinggi ---")
		fmt.Printf("Nama    : %s\n", T[idxTerbaik].nama)
		fmt.Printf("NIM     : %s\n", T[idxTerbaik].nim)
		fmt.Printf("IPK     : %.2f\n", T[idxTerbaik].ipk)
	} else {
		fmt.Println("Tidak ada data mahasiswa.")
	}
}
