package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var oleng bool
	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong:")
		fmt.Scan(&a, &b)
		if (a+b) > 150 || a < 0 || b < 0 {
			break
		}
		oleng = math.Abs(a-b) >= 9
		fmt.Println("Sepeda motor Pak Andi akan oleng:", oleng)
	}
	fmt.Print("Proses selesai")
}
