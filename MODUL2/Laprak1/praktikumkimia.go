package main

import "fmt"

func main() {
	var g1, g2, g3, g4 string
	var percobaan, true int

	true = 0

	for percobaan = 1; percobaan <= 5; percobaan++ {
		fmt.Print("Percobaan ke ", percobaan, ":")
		fmt.Scan(&g1, &g2, &g3, &g4)
		if g1 == "merah" && g2 == "kuning" && g3 == "hijau" && g4 == "ungu" {
			true++
		}
	}
	if true == 5 {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
