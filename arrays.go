package main

import "fmt"

func main() {
	// Arrays
	var arr [5]int
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	arr[3] = 400
	arr[4] = 900
	fmt.Printf("Array: %v\n", arr)
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	anotherArray := [4]int{448, 309, 798, 681}
	fmt.Printf("Another Array: %v\n", anotherArray)

	otherArray := [...]int{604, 932, 212}
	fmt.Printf("Other Array: %v\n", otherArray)
	fmt.Printf("Length of Other Array: %d\n", len(otherArray))
	fmt.Printf("Capacity of Other Array: %d\n", cap(otherArray))

	// Slices
	array := [...]int{235, 789, 10, 949, 339}
	fmt.Printf("Array: %v\n", array)

	slice1 := array[:] // slice of all elements
	fmt.Printf("Slice All elementes: %v\n", slice1)

	slice2 := array[2:] // slice from 3rd element to end
	fmt.Printf("Slice from 3rd element to end: %v\n", slice2)

	slice3 := array[:4] // slice first 4 elements
	fmt.Printf("Slice first 4 elements: %v\n", slice3)

	slice4 := array[1:3] // slice the 2nd and 3rd elements
	fmt.Printf("Slice the 2nd and 3rd elements: %v\n", slice4)

	//When we slice, the slice will point to the original array
	//that means that if we change the original array, the slice will change too

	// Slice length and capacity
	sliceMake := make([]int, 5) // Type of slice, length, capacity
	fmt.Printf("Slice Make: %v\n", sliceMake)
	fmt.Printf("Length of Slice Make: %d\n", len(sliceMake))
	fmt.Printf("Capacity of Slice Make: %d\n", cap(sliceMake))

	sliceMake2 := make([]int, 3, 10) // Type of slice, length, capacity
	fmt.Printf("SliceMake2: %v\n", sliceMake2)
	fmt.Printf("Length of SliceMake2: %d\n", len(sliceMake2))
	fmt.Printf("Capacity of SliceMake2: %d\n", cap(sliceMake2))

	// Slice append
	a := []int{}
	fmt.Printf("Slice a: %v\n", a)
	fmt.Printf("Length of Slice a: %d\n", len(a))
	fmt.Printf("Capacity of Slice a: %d\n", cap(a))

	a = append(a, 1)
	fmt.Printf("Slice a: %v\n", a)
	fmt.Printf("Length of Slice a: %d\n", len(a))
	fmt.Printf("Capacity of Slice a: %d\n", cap(a))

	// Go have the spread operator too, like JS, that will spread each element of the array
	b := []int{368, 350, 662, 62, 576}
	a = append(a, b...) // if we do append(a, b) it will fail because of type mismatch
	fmt.Printf("Slice a: %v\n", a)
	fmt.Printf("Length of Slice a: %d\n", len(a))
	fmt.Printf("Capacity of Slice a: %d\n", cap(a))
}
