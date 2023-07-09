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

	log.Println("----------------------------------------")
	x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y = x[:2]
	z = x[2:]
	log.Printf("x : %d, y: %d, z: %d", cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	log.Printf("x value: %v", x)
	log.Printf("y value: %v", y)
	log.Printf("z value: %v", z)

	log.Println("----------------------------------------")

	x1 = make([]int, 0, 5)
	x1 = append(x1, 1, 2, 3, 4)
	y1 = x1[:2:2]
	z1 := x1[2:4:4]
	log.Printf("y1 cap : %d, z1 cap: %d", cap(y1), cap(z1))
	y1 = append(y1, 30, 40, 50)
	z1 = append(z1, 70)
	log.Printf("y1 value: %v", y1)
	log.Printf("z1 value: %v", z1)

}
