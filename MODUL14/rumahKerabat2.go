package main

import "fmt"

const NMAX int = 1000

type tabInt [NMAX]int

func urutGanjil(arr *tabInt, n int) {
	var i, j, idxmin, temp int
	for i = 0; i < n-1; i++ {
		idxmin = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[idxmin] {
				idxmin = j
			}
		}
		temp = arr[i]
		arr[i] = arr[idxmin]
		arr[idxmin] = temp
	}
}

func urutGenap(arr *tabInt, n int) {
	var i, j, idxmin, temp int
	for i = 0; i < n-1; i++ {
		idxmin = i
		for j = i + 1; j < n; j++ {
			if arr[j] > arr[idxmin] {
				idxmin = j
			}
		}
		temp = arr[i]
		arr[i] = arr[idxmin]
		arr[idxmin] = temp
	}
}

func printData(ganjil tabInt, nGanjil int, genap tabInt, nGenap int) {
	var pertama bool
	var i int
	pertama = true
	for i = 0; i < nGanjil; i++ {
		if !pertama {
			fmt.Print(" ")
		}
		fmt.Print(ganjil[i])
		pertama = false
	}

	for i = 0; i < nGenap; i++ {
		if !pertama {
			fmt.Print(" ")
		}
		fmt.Print(genap[i])
		pertama = false
	}
}

func main() {
	var nDaerah, nGanjil, nGenap, m, x, i, j int
	var ganjil, genap tabInt
	fmt.Scan(&nDaerah)
	for i = 0; i < nDaerah; i++ {
		nGanjil = 0
		nGenap = 0
		fmt.Scan(&m)
		for j = 0; j < m; j++ {
			fmt.Scan(&x)
			if x%2 == 1 {
				ganjil[nGanjil] = x
				nGanjil++
			} else {
				genap[nGenap] = x
				nGenap++
			}
		}
		urutGanjil(&ganjil, nGanjil)
		urutGenap(&genap, nGenap)
		printData(ganjil, nGanjil, genap, nGenap)
	}
}
