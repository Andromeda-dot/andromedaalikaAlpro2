package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hitungToppingPizza(totalTopping int) int {
	var i int
	var countPizza int = 0
	var x, y, jarak float64

	for i = 1; i <= totalTopping; i++ {
		x = rand.Float64()
		y = rand.Float64()

		jarak = (x-0.5)*(x-0.5) + (y-0.5)*(y-0.5)

		if jarak <= 0.25 {
			countPizza = countPizza + 1
		}
	}
	return countPizza
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var inputTopping, toppingPizza int

	fmt.Print("Banyak Topping: ")
	fmt.Scan(&inputTopping)

	toppingPizza = hitungToppingPizza(inputTopping)
	fmt.Printf("Topping pada pizza: %d\n", toppingPizza)
}
