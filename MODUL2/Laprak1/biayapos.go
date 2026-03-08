package main

import "fmt"

func main() {
	var berat, totalberat, kg, sisa, jasakirim, biayatambahan, totalbiaya int
	fmt.Print("Berat parsel (gram):")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000
	totalberat = kg + sisa/1000

	jasakirim = kg * 10000
	biayatambahan = 0

	if totalberat > 10 {
		biayatambahan = 0
	} else if sisa >= 500 {
		biayatambahan = sisa * 5
	} else {
		biayatambahan = sisa * 15
	}
	totalbiaya = jasakirim + biayatambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp %d + Rp %d\n", jasakirim, biayatambahan)
	fmt.Printf("Total biaya: Rp %d\n", totalbiaya)
}
