package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

type lingkaran struct {
	titik_pusat titik
	radius      float64
}

func jarak(p, q titik) float64 {
	var dx, dy float64
	dx = q.x - p.x
	dy = q.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.titik_pusat, p) <= c.radius
}

func main() {
	var t1, t2, p titik
	var r1, r2 float64
	var l1, l2 lingkaran
	var lingkaran1, lingkaran2 bool

	fmt.Scan(&t1.x, &t1.y, &r1)
	fmt.Scan(&t2.x, &t2.y, &r2)
	fmt.Scan(&p.x, &p.y)

	l1 = lingkaran{titik_pusat: t1, radius: r1}
	l2 = lingkaran{titik_pusat: t2, radius: r2}

	lingkaran1 = didalam(l1, p)
	lingkaran2 = didalam(l2, p)

	if lingkaran1 && lingkaran2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if lingkaran1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if lingkaran2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan lingkaran 2")
	}
}
