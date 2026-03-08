package main

import "fmt"

func main() {
	var pita, bunga string
	var n int

	fmt.Print("bunga 1:")
	fmt.Scan(&bunga)

	for bunga != "SELESAI" {
		pita = pita + bunga + " - "
		n++
		fmt.Print("Bunga ", n+1, ":")
		fmt.Scan(&bunga)
	}
	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", n)
}
