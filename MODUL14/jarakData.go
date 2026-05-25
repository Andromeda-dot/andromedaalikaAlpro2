package main

import "fmt"

const NMAX int = 1000

type tabInt [NMAX]int

func pengurutanData(arr *tabInt, n int) {
	var i, j, key int
	for i = 1; i < n; i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func cekJarak(arr tabInt, n int) (bool, int) {
	var jarak, i int
	var jarakTetap bool
	jarakTetap = true
	if n < 2 {
		return true, 0
	}
	jarak = arr[1] - arr[0]
	for i = 2; i < n; i++ {
		if arr[i]-arr[i-1] != jarak {
			jarakTetap = false
		}
	}
	return jarakTetap, jarak
}

func main() {
	var data tabInt
	var jarak, n, x, i int
	var jarakTetap bool
	n = 0
	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		data[n] = x
		n++
	}
	pengurutanData(&data, n)
	for i = 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	jarakTetap, jarak = cekJarak(data, n)

	if jarakTetap {
		fmt.Println("Data berjarak ", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
