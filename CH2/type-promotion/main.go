package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 20.2
	// fmt.Print(x + y) error
	fmt.Print(float64(x) + y) // pass
}
