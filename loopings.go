package main

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		println(i)
	}

	// while loop
	i := 0
	for i < 5 {
		println(i)
		i++
	}

	// infinite loop
	for {
		println("Infinite loop")
		break
	}

	// for loop with range
	s := []int{1, 2, 3}
	for k, v := range s {
		println(k, v)
	}

	// for loop with range and ignore index
	for _, v := range s {
		println(v)
	}

	// for loop with range and ignore value
	for k, _ := range s {
		println(k)
	}

	// for loop with range and ignore both index and value
	for range s {
		println("Looping")
	}

	// Loops with label
outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			println(i, j)
			if i == 2 {
				break outer
			}
		}
	}
}
