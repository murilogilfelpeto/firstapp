package main

import "fmt"

func main() {
	// Create a map
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Printf("State populations: %v\n", statePopulations)

	// Create with make
	brazilianStates := make(map[string]int)
	brazilianStates = map[string]int{
		"SP": 39250017,
		"RJ": 27862596,
		"MG": 20612439,
		"BA": 19745289,
	}

	fmt.Printf("Brazilian states: %v\n", brazilianStates)

	// Add a new element
	statePopulations["Georgia"] = 10310371
	fmt.Printf("added State population: %v\n", statePopulations)

	// Delete an element
	delete(statePopulations, "Georgia")
	fmt.Printf("deleted State population: %v\n", statePopulations)

	// Get an element
	fmt.Printf("Get State population: %v\n", statePopulations["California"])

	// Get an element that doesn't exist
	fmt.Printf("Get State population that does not exist: %v\n", statePopulations["Georgia"])

	// Update an element
	statePopulations["California"] = 5
	fmt.Printf("updated State population: %v\n", statePopulations)

	// Verify if an element exists
	pop, ok := statePopulations["Georgia"]
	fmt.Printf("Verify if an element exists: %v, %v\n", pop, ok)
	pop, ok = statePopulations["California"]
	fmt.Printf("Verify if an element exists: %v, %v\n", pop, ok)

	// Get the length
	fmt.Printf("Get the length: %v\n", len(statePopulations))

	// Since Go is pass by reference, if you pass a map to a function and change it, it will change the original map
	sp := statePopulations
	delete(sp, "California")
	fmt.Printf("Original map: %v\n", statePopulations)
	fmt.Printf("Copy map: %v\n", sp)

	// If you want to pass a copy of the map, you need to create a new map and copy the values
	bs := make(map[string]int)
	for k, v := range brazilianStates {
		bs[k] = v
	}
	delete(bs, "SP")
	fmt.Printf("Original map: %v\n", brazilianStates)
	fmt.Printf("Copy map: %v\n", bs)

	// Explaining the difference between maps and structs
	// Maps are reference types
	// Maps are not comparable
	// Keys must be comparable
	// Use maps when you need to look up data by key
	// Use structs for more complex data structures
	// Examples of Structs:
	type Doctor struct {
		number     int
		actorName  string
		companions []string
	}
	doctorWho := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Printf("Struct: %v Type: %T\n", doctorWho, doctorWho)
	fmt.Printf("doctor number: %v\n", doctorWho.number)
	fmt.Printf("doctor actorName: %v\n", doctorWho.actorName)
	fmt.Printf("doctor companions: %v\n", doctorWho.companions)
	// Iterate over doctorWho companions
	for _, companion := range doctorWho.companions {
		fmt.Printf("doctor companion: %v\n", companion)
	}

	// Go does not have inheritance, but we can simulate it in the following way:
	type Animal struct {
		name   string
		origin string
	}
	type Bird struct {
		Animal
		speed  float32
		canFly bool
	}

	pidgeon := Bird{}
	pidgeon.name = "Roberto"
	pidgeon.origin = "street"
	pidgeon.canFly = true
	pidgeon.speed = 36

	fmt.Printf("Pidgeon: %v", pidgeon)
}
