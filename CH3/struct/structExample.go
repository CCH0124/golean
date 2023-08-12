package main

import (
	"log"
)

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var fred person
	log.Printf("fred: %v", fred)
	bob := person{}
	julia := person{
		"Julia",
		40,
		"cat",
	}
	beth := person{
		name: "Beth",
		age:  40,
	}

	bob.name = "Bob"
	log.Printf("bob: %v", bob)
	log.Printf("julia: %v", julia)
	log.Printf("beth name: %s", beth.name)
}
