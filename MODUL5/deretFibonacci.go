package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var suku_n, i int
	fmt.Print("n = ")
	fmt.Scan(&suku_n)

	for i = 0; i <= suku_n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}

}
