package main

import "fmt"

// Go supports pointers,
// allowing you to pass references to values and records within your program.

func main() {
	/*
		This code snippet declares a variable a of type int and assigns it a value of 42.
		It then declares a pointer variable b of type *int and assigns it the memory address of a.
		The code then prints the values of a and b, the memory addresses of a and b,
		and the value of a and the value stored at the memory address b points to.
		Finally, it updates the value of a to 27
		and prints the new value of a and the value stored at the memory address b points to.
	*/
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 13
	fmt.Println(a, *b)

	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	fmt.Printf("%v %p %p\n", x, y, z)

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	var otherStruct *myStruct
	fmt.Println(otherStruct)
	otherStruct = new(myStruct)
	otherStruct.foo = 99
	fmt.Println(otherStruct)
	fmt.Println(otherStruct.foo)
}

type myStruct struct {
	foo int
}
