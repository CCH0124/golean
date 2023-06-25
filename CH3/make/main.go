package main

import (
	"log"
)

func main() {

	var z = make([]int, 0, 3)
	z = append(z, 10)
	log.Printf("value: %v", z)
}
