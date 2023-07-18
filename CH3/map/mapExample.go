package main

import "log"

func main() {
	team := map[string][]string{
		"a": []string{"Itachi", "Madara"},
		"b": []string{"Kevin", "Bob"},
		"c": []string{"Kitty", "Kathy"},
	}

	log.Printf("team: %v", team)
	log.Println("----------------------------------------------------")

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	log.Printf("Orcas: %d", totalWins["Orcas"])
	log.Printf("Kittens: %d", totalWins["Kittens"])

	totalWins["Kittens"]++
	log.Printf("Kittens: %d", totalWins["Kittens"])
	totalWins["Lions"] = 3
	log.Printf("Lions: %d", totalWins["Lions"])

	log.Println("----------------------------------------------------")

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	log.Printf("hello v is %d, key is %t", v, ok)

	v, ok = m["golang"]
	log.Printf("golang v is %d, key is %t", v, ok)

	delete(m, "hello")

	v, ok = m["hello"]
	log.Printf("hello v is %d, key is %t", v, ok)
}
