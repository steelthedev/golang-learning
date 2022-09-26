package main

import "fmt"

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayGoodbye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func main() {
	sayGreeting("Steel")
	sayGoodbye("Steel")

}
