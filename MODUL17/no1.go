package main

import "fmt"

func main() {
	var bilangan, total, rata2 float64
	var hitung int

	total = 0
	hitung = 0

	for {
		fmt.Scan(&bilangan)
		if bilangan == 9999 {
			break
		}
		total += bilangan
		hitung++
	}

	if hitung > 0 {
		rata2 = total / float64(hitung)
		fmt.Printf("Jumlah bilangan yang valid: %d\n", hitung)
		fmt.Printf("Total penjumlahan: %.2f\n", total)
		fmt.Printf("Rata-Rata: %.2f\n", rata2)
	} else {
		fmt.Println("Tidak ada bilangan valid!")
	}
}
