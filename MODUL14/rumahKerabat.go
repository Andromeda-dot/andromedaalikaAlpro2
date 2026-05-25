package main

import "fmt"

const NMAX int = 1000

type tabInt [NMAX]int

func urutanRumah(arr *tabInt, n int) {
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

func printData(arr tabInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print(arr[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print()
}

func main() {
	var nDaerah, m, i, j int
	var rumah tabInt

	fmt.Scan(&nDaerah)

	for i = 0; i < nDaerah; i++ {
		fmt.Scan(&m)
		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		urutanRumah(&rumah, m)
		printData(rumah, m)
	}
}
