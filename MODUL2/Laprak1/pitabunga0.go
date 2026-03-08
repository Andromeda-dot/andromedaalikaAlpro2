package main

import "fmt"

func main() {
	var pita, bunga string
	var n, i int

	fmt.Print("N: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Print("bunga: ")
		fmt.Scan(&bunga)
		pita = pita + bunga + " - "
	}
	fmt.Println("Pita: ", pita)
}
