package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s- save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("What is the item name?: ", reader)
		price, _ := getInput("price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("Items added - ", name, p)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file -", b.name)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("Tip added - ", tip)
		promptOptions(b)
	default:
		fmt.Println("Not a valid option")
		promptOptions(b)

	}

}
func main() {
	mybill := createBill()
	promptOptions(mybill)

}
