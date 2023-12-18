package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// If statement
	x := 10
	if x > 5 {
		println("x is greater than 5")
	} else {
		println("x is less than 5")
	}

	// Logic operators
	// && and
	// || or
	// ! not
	// == equal
	// != not equal
	// < less than
	// > greater than
	// <= less than or equal
	// >= greater than or equal

	// Switch statement
	y := 40
	switch y {
	case 10:
		fmt.Printf("y is equal to 10\n")
	case 20:
		fmt.Printf("y is equal to 20\n")
	default:
		fmt.Printf("y is not equal to 10 or 20\n")
	}

	numberGreaterThanFive(x)

	// Grouped case
	number := rand.Intn(10)
	switch number {
	case 0, 2, 4, 6, 8:
		fmt.Printf("We got an even number %v", number)
	case 1, 3, 5, 7, 9:
		fmt.Printf("We got an odd number %v", number)
	default:
		fmt.Printf("We got an unexpected number %v", number)
	}

	// Logical operators at case
	switch {
	case number > 5:
		fmt.Printf("We got a number greater than 5 %v\n", number)
	case number < 5:
		fmt.Printf("We got a number less than 5 %v\n", number)
	default:
		fmt.Printf("We got a number equal to 5 %v\n", number)
	}

	// Type switch
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Printf("%v is an integer \n", i)
	case float64:
		fmt.Printf("%v is a float64 \n", i)
	case string:
		fmt.Printf("%v is a string \n", i)
	default:
		fmt.Printf("%v is another type \n")
	}
}

func numberGreaterThanFive(number int) bool {
	if number > 5 {
		fmt.Println("x is greater than 5")
		return true
	} else {
		fmt.Println("x is less than 5")
		return false
	}
}
