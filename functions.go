package main

import "fmt"

func main() {
	// Call the function
	sayHello()
	saySomething("Alan Mehmood")
	for i := 0; i < 10; i++ {
		countMessages("John Doe", i+1)
	}
	total := sum(40, 282, 782, 801)
	fmt.Println("Total:", total)

	result, err := divide(1331, 2)
	fmt.Printf("Result: %.2f\n", result)
	fmt.Printf("Error: %v\n", err)

	result, err = divide(1331, 0)
	fmt.Printf("Result: %.2f\n", result)
	fmt.Printf("Error: %v\n", err)

	func() {
		fmt.Println("Hello from anonymous function!")
	}()

	// for anonymous functions, the best practice is always provide parameters
	// since they can access all variables outside the function, you can have some issues
	// when dealing with async code

	for i := 0; i < 10; i++ {
		func(i int) {
			fmt.Println("Hello from anonymous function!", i)
		}(i)
	}

	// we can have variables as functions
	f := func() {
		fmt.Println("Hello from variable function!")
	}
	f()

	var division func(float64, float64) (float64, error)
	division = func(n1 float64, n2 float64) (float64, error) {
		if n2 == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}
		return n1 / n2, nil
	}
	result, err = division(1331, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result: %.2f\n", result)

	greeter := greeter{
		greeting: "Hello",
		name:     "John",
	}
	greeter.greet()
}

type greeter struct {
	greeting string
	name     string
}

// This a method for Greeter struct
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func divide(number1, number2 float64) (float64, error) {
	if number2 == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return number1 / number2, nil
}

func sum(numbers ...int) int {
	fmt.Println(numbers, " ")
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total

}

func saySomething(msg string) {
	fmt.Printf("Hello %s!\n", msg)
}

func countMessages(msg string, i int) {
	fmt.Printf("Hello %s! %d° time\n", msg, i)
}

// Define the function
func sayHello() {
	fmt.Println("Hello!")
}
