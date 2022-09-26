package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayGoodbye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}
func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func cycleArea(r float64) float64 {
	return math.Pi * r * r

}

func main() {
	sayGreeting("Steel")
	sayGoodbye("Steel")
	cycleName([]string{"cloud", "thunder", "snow"}, sayGreeting)

	a1 := cycleArea(4.5)
	a2 := cycleArea(85)

	fmt.Println(a1, a2)
	fmt.Printf("cycle 1 is %0.3f and 2 is %0.3f", a1, a2)
}
