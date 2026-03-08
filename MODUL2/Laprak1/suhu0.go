package main

import "fmt"

func main() {
	var celcius, fahreinheit int
	fmt.Print("Derajat Celcius: ")
	fmt.Scan(&celcius)

	fahreinheit = (celcius * 9 / 5) + 32

	fmt.Println("Derajat Fahreinheit:", fahreinheit)
}
