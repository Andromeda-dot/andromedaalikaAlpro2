package main

import (
	"fmt"
)

// Definisi struct mahasiswa sesuai modul
type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

// Definisi array of struct dengan kapasitas 2023
type arrMhs [2023]mahasiswa

// Fungsi Sequential Search: mencari berdasarkan NAMA
func SeqSearch_3(T arrMhs, n int, X string) int {
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if T[j].nama == X {
			found = j
		}
		j = j + 1
	}
	return found
}

// Fungsi Binary Search: mencari berdasarkan NIM (syarat: array HARUS terurut berdasarkan NIM)
func BinarySearch_3(T arrMhs, n int, X string) int {
	var found int = -1
	var med int
	var kr int = 0
	var kn int = n - 1

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if X < T[med].nim {
			kn = med - 1
		} else if X > T[med].nim {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

func main() {
	var dataMahasiswa arrMhs
	var n int = 4

	// Mengisi data dummy
	// PERHATIKAN: Data sengaja diurutkan dari NIM terkecil ke terbesar
	dataMahasiswa[0] = mahasiswa{nama: "Andi", nim: "1301210001", kelas: "IF-45-01", jurusan: "Informatika", ipk: 3.8}
	dataMahasiswa[1] = mahasiswa{nama: "Budi", nim: "1301210005", kelas: "IF-45-01", jurusan: "Informatika", ipk: 3.5}
	dataMahasiswa[2] = mahasiswa{nama: "Citra", nim: "1301210010", kelas: "IF-45-02", jurusan: "Informatika", ipk: 3.9}
	dataMahasiswa[3] = mahasiswa{nama: "Ren", nim: "1301210020", kelas: "IF-45-02", jurusan: "Informatika", ipk: 4.0}

	fmt.Println("=== Demo 1: Sequential Search (Cari Nama) ===")
	targetNama := "Citra"
	idxNama := SeqSearch_3(dataMahasiswa, n, targetNama)

	if idxNama != -1 {
		fmt.Printf("Ketemu! %s ada di indeks ke-%d\n", targetNama, idxNama)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}

	fmt.Println("\n=== Demo 2: Binary Search (Cari NIM) ===")
	targetNIM := "1301210020"
	idxNIM := BinarySearch_3(dataMahasiswa, n, targetNIM)

	if idxNIM != -1 {
		fmt.Printf("Ketemu! NIM %s adalah milik %s di indeks ke-%d\n", targetNIM, dataMahasiswa[idxNIM].nama, idxNIM)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}
