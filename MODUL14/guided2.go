package main

import "fmt"

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	fmt.Println(arr)
}
