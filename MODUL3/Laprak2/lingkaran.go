package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	var jarak float64
	jarak = math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
	return jarak
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, kx, ky float64
	var lingkaran1, lingkaran2 bool

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&kx, &ky)

	lingkaran1 = didalam(cx1, cy1, r1, kx, ky)
	lingkaran2 = didalam(cx2, cy2, r2, kx, ky)

	if lingkaran1 && lingkaran2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if lingkaran1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if lingkaran2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan lingkaran 2")
	}
}
