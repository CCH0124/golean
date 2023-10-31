package main

import "log"

func main() {
	x := 10
	if x > 5 {
		log.Println(x)
		x := 5
		log.Println(x)
	}
	log.Println(x)
}
