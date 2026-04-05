package main

import "fmt"

func faktorial(n int, hasil *int) {
	var i int
	*hasil = 1
	for i = 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFaktorial, nrFaktorial int
	faktorial(n, &nFaktorial)
	faktorial(n-r, &nrFaktorial)

	*hasil = nFaktorial / nrFaktorial
}

func combination(n, r int, hasil *int) {
	var nFaktorial, rFaktorial, nrFaktorial int

	faktorial(n, &nFaktorial)
	faktorial(r, &rFaktorial)
	faktorial(n-r, &nrFaktorial)

	*hasil = nFaktorial / (rFaktorial * nrFaktorial)
}

func main() {
	var a, b, c, d int
	var permutasi, kombinasi int
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &permutasi)
	combination(a, c, &kombinasi)
	fmt.Println(permutasi, kombinasi)

	permutation(b, d, &permutasi)
	combination(b, d, &kombinasi)
	fmt.Println(permutasi, kombinasi)
}
