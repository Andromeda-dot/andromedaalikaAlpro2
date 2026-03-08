package main

import "fmt"

func main() {
	var jari2 int
	var volume, luas float64
	fmt.Print("Jari-jari: ")
	fmt.Scan(&jari2)

	volume = (4.0 / 3.0) * 3.1415926535 * float64(jari2) * float64(jari2) * float64(jari2)
	luas = 4 * 3.1415926535 * float64(jari2) * float64(jari2)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", jari2, volume, luas)
}
