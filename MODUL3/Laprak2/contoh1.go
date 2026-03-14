package main

import "fmt"

func faktorial(n int) int {
	var hasil, i int
	hasil = 1
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func main() {
	var a, b int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println("Hasil: ", permutasi(a, b))
	} else {
		fmt.Println("Hasil: ", permutasi(b, a))
	}
}
