package main

import "fmt"

func cetakBilangan(n int) {
	if n == 1 {
		fmt.Printf("%d ", n)
		return
	}
	fmt.Printf("%d ", n)
	cetakBilangan(n - 1)
	fmt.Printf("%d ", n)
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	cetakBilangan(bilangan)
}
