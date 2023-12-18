package main

import "fmt"

const value int = 50

// How to declare enums, can be done in a Block or normal way
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

func main() {

	//To declare constants we need to follow the camelCase equal variables
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// We can modify a constant value when we declare one at package level
	//and other inside the function
	fmt.Printf("before changing its value: %v, %T\n", value, value)
	const value int16 = 13
	fmt.Printf("after changing its value: %v, %T\n", value, value)

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)

}
