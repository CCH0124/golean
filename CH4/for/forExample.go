package main

import "log"

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	uniqName := map[string]bool{"a": true, "b": true, "c": true}
	for k := range uniqName {
		log.Println(k)
	}

	evenVals := []int{2, 4, 6, 8}
	for _, v := range evenVals {
		v *= 2
	}
	log.Println(evenVals)
}
