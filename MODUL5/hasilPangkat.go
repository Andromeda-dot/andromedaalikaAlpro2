package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pangkat(x, y-1)
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(pangkat(a, b))
}
