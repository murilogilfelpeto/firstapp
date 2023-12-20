package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

func main() {
	go logger()
	logCh <- logEntry{time.Now(), "info", "App is starting"}

	channel := make(chan int)
	wg.Add(2)
	go func() {
		i := <-channel // Receive message from the channel
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		channel <- 42 // Send message to the channel
		wg.Done()
	}()
	wg.Wait()

	for j := 0; j < 10; j++ {
		wg.Add(2)
		go func() {
			number := <-channel
			fmt.Println(number)
			wg.Done()
		}()
		j := j
		go func() {
			channel <- j
			wg.Done()
		}()
	}
	wg.Wait()

	wg.Add(2)
	// Receive only go routine
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(channel)

	// Send only go routine
	go func(ch chan<- int) {
		ch <- 234
		wg.Done()
	}(channel)
	wg.Wait()

	// We can provide a buffer to channel
	bufferedChannel := make(chan int, 50)
	wg.Add(2)
	// Receive only go routine
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Printf("Received: %d\n", i)
			} else {
				break
			}
		}
		wg.Done()
	}(bufferedChannel)

	// Send only go routine
	go func(ch chan<- int) {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(bufferedChannel)
	wg.Wait()

	logCh <- logEntry{time.Now(), "logInfo", "App is shutting down"}
	doneCh <- struct{}{}
}

type logEntry struct {
	Time     time.Time
	Severity string
	Message  string
}

func logger() {
	for {
		select {
		case <-doneCh:
			break
		case log := <-logCh:
			fmt.Printf("%v - [%v] - %v\n", log.Time.Format("2006-01-02T15:04:05"), log.Severity, log.Message)
		}
	}
}
