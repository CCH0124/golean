package main

import (
	"log"
)

type firstPerson struct {
	name string
	age  int
}

func main() {
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	g = f
	log.Printf("g: %v", g)
	log.Printf(" f == g : %v", f == g)
}
