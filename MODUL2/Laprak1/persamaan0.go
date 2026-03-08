package main

import "fmt"

func main() {
	var k, pembilang, penyebut, fk float64
	fmt.Print("Nilai K: ")
	fmt.Scan(&k)

	pembilang = ((4 * k) + 2) * ((4 * k) + 2)
	penyebut = ((4 * k) + 1) * ((4 * k) + 3)
	fk = pembilang / penyebut
	fmt.Printf("Nilai f(k): %.10f\n", fk)
}
