package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup = sync.WaitGroup{}
var counterRaceCondition = 0
var counterMutex = 0
var mutex = sync.RWMutex{}

func main() {
	/*
		creates a new goroutine (a lightweight thread) to print "Hello Thread" to the console.
		It uses a sync.WaitGroup to wait for the goroutine to finish before exiting the program.
	*/
	waitGroup.Add(1) // add 1 to the wait group
	go func() {
		var msg = "Hello Thread"
		fmt.Println(msg)
		waitGroup.Done() // subtract 1 from the wait group
	}()
	waitGroup.Wait() // block the main thread until the wait group counterRaceCondition is 0

	// This will create a race condition
	for i := 0; i < 20; i++ {
		waitGroup.Add(2)
		go messageRaceCondition()
		go incrementRaceCondition()
	}
	waitGroup.Wait()

	runtime.GOMAXPROCS(200)
	for i := 0; i < 20; i++ {
		waitGroup.Add(2)
		mutex.RLock()
		go messageMutex()
		mutex.Lock()
		go incrementMutex()
	}
	waitGroup.Wait()
}

func messageRaceCondition() {
	fmt.Printf("Hello from race condition, Thread#%v\n", counterRaceCondition)
	waitGroup.Done()
}

func incrementRaceCondition() {
	counterRaceCondition++
	waitGroup.Done()
}

func messageMutex() {

	fmt.Printf("Hell from mutex, Thread #%v\n", counterMutex)
	mutex.RUnlock()
	waitGroup.Done()
}

func incrementMutex() {
	counterMutex++
	mutex.Unlock()
	waitGroup.Done()
}
