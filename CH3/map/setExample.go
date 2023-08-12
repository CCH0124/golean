package main

import (
	"log"
)

func main() {
	intSet := map[int]bool{}

	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}

	log.Printf("vals len: %d, intSet len: %d", len(vals), len(intSet))
	log.Print(intSet[5])
	log.Print(intSet[500])
	if intSet[100] {
		log.Println("not in the set.")
	}
}
