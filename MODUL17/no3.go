package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hitungTetesanHujan(totalTetesan int) (int, int, int, int) {
	var x, y float64
	var countA int = 0
	var countB int = 0
	var countC int = 0
	var countD int = 0
	var i int

	for i = 1; i <= totalTetesan; i++ {
		x = rand.Float64()
		y = rand.Float64()

		if x < 0.5 && y < 0.5 {
			countA = countA + 1
		} else if x >= 0.5 && y < 0.5 {
			countB = countB + 1
		} else if x >= 0.5 && y >= 0.5 {
			countC = countC + 1
		} else if x < 0.5 && y >= 0.5 {
			countD = countD + 1
		}
	}
	return countA, countB, countC, countD
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var input int
	fmt.Print("Masukkan banyaknya tetesan air hujan: ")
	fmt.Scan(&input)

	var tetesanA int = 0
	var tetesanB int = 0
	var tetesanC int = 0
	var tetesanD int = 0

	tetesanA, tetesanB, tetesanC, tetesanD = hitungTetesanHujan(input)

	var faktorKonversi float64 = 0.0001

	var curahHujanA float64 = float64(tetesanA) * faktorKonversi
	var curahHujanB float64 = float64(tetesanB) * faktorKonversi
	var curahHujanC float64 = float64(tetesanC) * faktorKonversi
	var curahHujanD float64 = float64(tetesanD) * faktorKonversi

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", curahHujanA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", curahHujanB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", curahHujanC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", curahHujanD)
}
