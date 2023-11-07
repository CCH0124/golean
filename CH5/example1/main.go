package main

import "log"

func main() {
	res := div(5, 2)
	log.Println(res)
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
