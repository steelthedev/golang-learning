package main

import "fmt"

func main() {
	age := 60

	if age <= 70 {
		fmt.Println(age)
	} else if age < 60 {
		fmt.Println("Age is lesser than 60")
	} else {
		fmt.Println("I dont want to be invloved in this conditional statement")
	}

	names := []string{"aisha", "gbenro", "blessing", "adeola", "iyanu"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at position", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}
