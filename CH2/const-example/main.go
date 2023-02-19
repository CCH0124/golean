package main

import "fmt"

const x int64 = 10
const (
	id   = "id"
	name = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "world"

	fmt.Println(x)
	fmt.Println(y)

}
