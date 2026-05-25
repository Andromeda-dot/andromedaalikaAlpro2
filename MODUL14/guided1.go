package main

import "fmt"

type arrInt [4321]int

func main() {
	var i int = 1
	var n, idx_min, j, t int
	var a arrInt

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i

		for j < n {
			if a[idx_min] > a[j] {
				idx_min = j
			}
			j = j + 1
		}

		t = a[idx_min]
		a[idx_min] = a[i-1]
		a[i-1] = t

		i = i + 1
	}

	for i = 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
}
