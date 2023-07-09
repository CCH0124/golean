package main

import (
	"log"
)

func main() {
	var s string = "Hello String"
	var b byte = s[6]
	log.Printf("%b", b)
	var s2 string = s[4:7]
	// var s3 string = s[:5]
	// var s4 string = s[6:]

	var a rune = 'x'
	s = string(a)
	var a1 rune = 'y'
	s2 = string(a1)
	log.Printf("s: %s", s)
	log.Printf("s2: %s", s2)

	s = "Hello, 🥱"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	log.Printf("bs: %v", bs)
	log.Printf("rs: %v", rs)

}
