package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

var tab tabel
var m int

func isiArray(t *tabel, n *int) {
	var input rune
	fmt.Scanf("%c", &input)
	*n = 0
	for input != '.' && *n < NMAX {
		t[*n] = input
		*n++
		fmt.Scanf("%c", &input)
	}
}

func cetakArray(t tabel, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	var temp rune
	var i int
	for i = 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var i int
	for i = 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Karakter: ")
	isiArray(&tab, &m)

	balikkanArray(&tab, m)

	fmt.Print("Isi array setelah dibalik: ")
	cetakArray(tab, m)

	if palindrom(tab, m) {
		fmt.Print("Palindrom: true")
	} else {
		fmt.Print("Palindrom: false")
	}
}
