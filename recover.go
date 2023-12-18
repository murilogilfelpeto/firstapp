package main

import "fmt"

// It prints "Start", calls the panicker function, and then prints "End".
func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

// panicker is a Go function that demonstrates panic and recover.
// It does not have any parameters or return types.
func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	panic("Something bad happened")
}
