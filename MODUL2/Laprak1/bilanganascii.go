package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var ch1, ch2, ch3 byte

	fmt.Scanln(&a, &b, &c, &d, &e)
	fmt.Scanf("%c%c%c\n", &ch1, &ch2, &ch3)

	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)
	fmt.Printf("%c%c%c\n", ch1+1, ch2+1, ch3+1)
}
