package main

import "fmt"

func main() {
	var celcius, fahreinheit, reamur, kelvin int
	fmt.Print("Derajat Celcius: ")
	fmt.Scan(&celcius)

	reamur = celcius * 4 / 5
	fahreinheit = (celcius * 9 / 5) + 32
	kelvin = (fahreinheit + 460) * 5 / 9

	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahreinheit: ", fahreinheit)
	fmt.Println("Derajat Kelvin: ", kelvin)
}
