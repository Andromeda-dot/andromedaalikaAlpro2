package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	ekslempar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	var i int
	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].ekslempar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var idxfav, i int
	idxfav = 0
	for i = 0; i < n; i++ {
		if pustaka[i].rating > pustaka[idxfav].rating {
			idxfav = i
		}
	}
	fmt.Printf("%s, %s, %s, %d\n",
		pustaka[idxfav].judul, pustaka[idxfav].penulis, pustaka[idxfav].penerbit, pustaka[idxfav].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n *int) {
	var i, j int
	var key Buku
	for i = 1; i < *n; i++ {
		key = pustaka[i]
		j = i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var batas, i int
	batas = 5
	if n < 5 {
		batas = n
	}
	for i = 0; i < batas; i++ {
		fmt.Print(pustaka[i].judul, " ")
	}
	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, ketemu, median int
	var b Buku
	kiri = 0
	kanan = n - 1
	ketemu = -1
	for kiri <= kanan && ketemu == -1 {
		median = (kiri + kanan) / 2
		if pustaka[median].rating == r {
			ketemu = median
		} else if pustaka[median].rating < r {
			kanan = median - 1
		} else {
			kiri = median + 1
		}
	}
	if ketemu != -1 {
		b = pustaka[ketemu]
		fmt.Printf("%s %s %s %d %d %d\n", b.judul, b.penulis, b.penerbit,
			b.tahun, b.ekslempar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka, ratingDicari int

	DaftarkanBuku(&pustaka, &nPustaka)
	fmt.Print("Masukkan rating buku yang dicari: ")
	fmt.Scan(&ratingDicari)

	CetakTerfavorit(pustaka, nPustaka)

	UrutBuku(&pustaka, &nPustaka)

	Cetak5Terbaru(pustaka, nPustaka)

	CariBuku(pustaka, nPustaka, ratingDicari)
}
