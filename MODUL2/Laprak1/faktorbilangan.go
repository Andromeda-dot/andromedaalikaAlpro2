package main

import "fmt"

func main() {
	var bilangan, faktor, f int
	var prima bool
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)

	faktor = 0
	fmt.Print("Faktor: ")
	for f = 1; f <= bilangan; f++ {
		if bilangan%f == 0 {
			fmt.Print(f, " ")
			faktor++
		}
	}
	if faktor == 2 {
		prima = true
	} else {
		prima = false
	}
	fmt.Println()
	fmt.Println("Prima: ", prima)
}
