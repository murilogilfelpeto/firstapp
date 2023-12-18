package main

import (
	"fmt"
	"strconv"
)

// We can declare variables at package level, but we need to declare them telling its type
var packageMessage string = "Hello, package variable!"

// We can declare multiple variables at once just using a var block
var (
	text      string = "var 1"
	otherText        = "var 2"
)

func main() {
	// Declare a variable of type string with the value "Hello, World!" and print it.
	var message string = "Hello, variable String!"
	fmt.Println(message)

	// Go can discover by itself the type of variable, so we can write this:
	message2 := "go discover this type!"
	fmt.Println(message2)

	fmt.Println(packageMessage)

	//we can declare the same variable twice, but they need to be in different scopes, like this:
	packageMessage = "Override package message variable"
	fmt.Println(packageMessage)

	// In order to convert a number to String, we must import a package called strconv

	//Example without strconv
	var number int = 42
	var numberString string = string(number)
	fmt.Println("Without strconv")
	fmt.Printf("%v, %T\n", numberString, numberString)

	var usingStrConv string = strconv.Itoa(number)
	fmt.Println("Using strconv")
	fmt.Printf("%v, %T\n", usingStrConv, usingStrConv)
}
