package main

import (
	"log"
	"math/rand"
)

func main() {
	if n := rand.Intn(10); n == 0 {
		log.Println("too low")
	} else if n > 5 {
		log.Printf("too big %d", n)
	} else {
		log.Printf("good number %d", n)
	}
}
