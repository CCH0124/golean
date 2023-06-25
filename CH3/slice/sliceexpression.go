package main

import (
	"log"
)

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	log.Printf("y value: %v", y)
	z := x[1:]
	log.Printf("z value: %v", z)
	d := x[1:3]
	log.Printf("d value: %v", d)
	a := x[:]
	log.Printf("a value: %v", a)

	x[1] = 20
	y[0] = 10
	z[2] = 30
	log.Printf("x value: %v", x)
	log.Printf("y value: %v", y)
	log.Printf("z value: %v", z)

	x1 := []int{1, 2, 3, 4}
	y1 := x1[:2]
	log.Printf("x1 Cap: %d, y1 Cap: %d", cap(x1), cap(y1))

	y1 = append(y1, 30)
	log.Printf("x1 value: %v", x1)
	log.Printf("y1 value: %v", y1)

}
