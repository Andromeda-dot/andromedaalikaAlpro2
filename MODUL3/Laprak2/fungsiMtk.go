package main

import "fmt"

func fX(x int) int {
	var f int
	f = x * x
	return f
}

func gX(x int) int {
	var g int
	g = x - 2
	return g
}

func hX(x int) int {
	var h int
	h = x + 1
	return h
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(fX(gX(hX(a))))
	fmt.Println(gX(hX(fX(b))))
	fmt.Println(hX(fX(gX(c))))
}
