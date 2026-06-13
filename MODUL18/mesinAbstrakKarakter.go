package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4.a) Operasi Dasar Mesin Karakter
var pitaKarakter string
var indeksPita int
var currentChar byte
var EOP bool

// Prosedur start
func start(input string) {
	pitaKarakter = input
	indeksPita = 0
	EOP = false

	if len(pitaKarakter) > 0 {
		currentChar = pitaKarakter[indeksPita]
		if currentChar == '.' {
			EOP = true
		}
	} else {
		EOP = true
	}
}

// Prosedur maju
func maju() {
	if !EOP {
		indeksPita++
		if indeksPita < len(pitaKarakter) {
			currentChar = pitaKarakter[indeksPita]
			if currentChar == '.' {
				EOP = true
			}
		} else {
			EOP = true
		}
	}
}

// Fungsi eop
func eop() bool {
	return EOP
}

// Fungsi cc
func cc() byte {
	return currentChar
}

// 4.b) Algoritma
// Point a & b = Membaca seluruh karakter dan menghitung berapa banyak karakter terbaca
func hitungSemuaKarakter() int {
	var hitung int = 0

	for !eop() {
		hitung++
		maju()
	}
	return hitung
}

// Point c = Menghitung banyak huruf "A" yang terbaca
func hitungHurufA() int {
	var hitungA int = 0

	for !eop() {
		if cc() == 'A' || cc() == 'a' {
			hitungA++
		}
		maju()
	}
	return hitungA
}

// Point d = Menghitung frekuensi kemunculan huruf "A"
func hitungFrekuensiA() float64 {
	var hitungA int = 0
	var hitungTotal int = 0

	for !eop() {
		hitungTotal++
		if cc() == 'A' || cc() == 'a' {
			hitungA++
		}
		maju()
	}
	if hitungTotal == 0 {
		return 0.0
	}
	return float64(hitungA) / float64(hitungTotal)
}

// Point e = Menghitung ada berapa kata "LE" (pasangan berurutan 'L' dan 'E')
func hitungKataLE() int {
	var hitungLE int = 0
	var lastChar byte

	if !eop() {
		lastChar = cc()
		maju()
	}

	for !eop() {
		if (lastChar == 'L' || lastChar == 'l') && (cc() == 'E' || cc() == 'e') {
			hitungLE++
		}
		lastChar = cc()
		maju()
	}
	return hitungLE
}

func main() {
	var pembaca *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var teksKalimat string
	var totalA, totalLE int

	fmt.Print("Masukkan teks: ")

	if pembaca.Scan() {
		teksKalimat = pembaca.Text()

		fmt.Println("Hasil analisis: ")
		start(teksKalimat)
		totalA = hitungHurufA()
		fmt.Printf("Jumlah huruf 'A' yang terbaca: %d\n", totalA)

		start(teksKalimat)
		totalLE = hitungKataLE()
		fmt.Printf("Jumlah kata 'LE' yang terbaca: %d\n", totalLE)
	}
}
