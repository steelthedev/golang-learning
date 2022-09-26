package main

import "fmt"

func Menu() {
	menu := map[string]float64{
		"soup":    499,
		"pie":     799,
		"salad":   699,
		"chicken": 355,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	//looping maps

	for key, value := range menu {
		println(key, "-", value)
	}
	// ints as int type

	phoneBook := map[int]string{
		22: "Akinwumi",
		24: "Adeola",
		25: "Alebiosu",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[22])

	phoneBook[22] = "Ayomiposi"

	fmt.Println(phoneBook)
}
