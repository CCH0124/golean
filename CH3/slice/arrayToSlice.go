package main

import (
	"log"
)

func main() {
	x := [4]int{1, 2, 3, 4}
	y := x[:2]
	z := x[2:]
	log.Printf("x: %v", x)
	log.Printf("y: %v", y)
	log.Printf("z: %v", z)
	log.Println("-------------------copy----------------")
	x1 := []int{1, 2, 3, 4}
	y1 := make([]int, 4)
	num := copy(y1, x1)
	log.Printf("y: %v", y1)
	log.Printf("num length: %v", num)
}
