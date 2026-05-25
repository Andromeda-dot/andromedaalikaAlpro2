package main

import "fmt"

const NMAX int = 1000000

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

func median(arr tabInt, n int) int {
	var tengah int
	if n%2 == 1 {
		tengah = n / 2
		return arr[tengah]
	}
	tengah = n / 2
	return (arr[tengah-1] + arr[tengah]) / 2
}

func main() {
	var data tabInt
	var n, a int
	n = 0
	for {
		fmt.Scan(&a)
		if a == -5313 {
			break
		}
		if a == 0 {
			pengurutanData(&data, n)
			fmt.Println(median(data, n))
		} else {
			data[n] = a
			n++
		}
	}
}
