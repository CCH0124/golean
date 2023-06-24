package main

import (
	"log"
)

func main() {
	var x []int

	log.Printf("len: %d", len(x))

	// 至少兩個參數
	x = append(x, 10)

	log.Printf("x: %v", x)

	x = append(x, 11, 12, 13)

	log.Printf("x: %v", x)

	slice2 := []int{20, 21, 22, 23}
	// slice2... 附加到 x slice
	x = append(x, slice2...)

	log.Printf("x: %v", x)

	log.Println("===========================CAP================================")

	var z []int
	log.Printf("value: %v, len: %d, cap: %d", z, len(z), cap(z))
	z = append(z, 10)
	log.Printf("value: %v, len: %d, cap: %d", z, len(z), cap(z))
	z = append(z, 20)
	log.Printf("value: %v, len: %d, cap: %d", z, len(z), cap(z))
	z = append(z, 30)
	log.Printf("value: %v, len: %d, cap: %d", z, len(z), cap(z))
	z = append(z, 40)
	log.Printf("value: %v, len: %d, cap: %d", z, len(z), cap(z))
	z = append(z, 50)
	log.Printf("value: %v, len: %d, cap: %d", z, len(z), cap(z))
}

// 2023/06/24 15:37:04 len: 0
// 2023/06/24 15:37:04 x: [10]
// 2023/06/24 15:37:04 x: [10 11 12 13]
// 2023/06/24 15:37:04 x: [10 11 12 13 20 21 22 23]
// 2023/06/24 15:37:04 ===========================CAP================================
// 2023/06/24 15:37:04 value: [], len: 0, cap: 0
// 2023/06/24 15:37:04 value: [10], len: 1, cap: 1
// 2023/06/24 15:37:04 value: [10 20], len: 2, cap: 2
// 2023/06/24 15:37:04 value: [10 20 30], len: 3, cap: 4
// 2023/06/24 15:37:04 value: [10 20 30 40], len: 4, cap: 4
// 2023/06/24 15:37:04 value: [10 20 30 40 50], len: 5, cap: 8
