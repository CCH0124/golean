package main

import "log"

func main() {
	log.Println(addTo(3))
	log.Println(addTo(1, 2))
}

func addTo(base int, vals ...int) []int {
	res := make([]int, 0, len(vals))
	for _, v := range vals {
		res = append(res, base+v)
	}

	return res
}
