package main

import "fmt"

func main() {
	var bilangan, faktor, f int
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
}
