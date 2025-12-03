package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	This is the demonstration for GO routines (lightweight threads)
	You will learn how the multi-thread works here
*/

type Result struct {
	value string
	err   error
}

func Worker(url string, wg *sync.WaitGroup, resultChannel chan Result) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)

	// Proper formatting for printing
	fmt.Printf("Url: %v\n", url)

	resultChannel <- Result{
		value: url,
		err:   nil,
	}
}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	resultChannel := make(chan Result, 3) // buffered to 3 since we spawn 3 workers

	/*
		GO Routine life cycle
		fan out --> main thread to worker (lightweight threads)
		fan in  --> worker thread to main thread
		main --> worker --> worker -> channel --> channel -> main thread
	*/

	wg.Add(3)
	go Worker("don.jpg", &wg, resultChannel)
	go Worker("don1.jpg", &wg, resultChannel)
	go Worker("don2.jpg", &wg, resultChannel)

	// wait for all workers to finish sending
	wg.Wait()

	// close channel after all sends are done
	close(resultChannel)

	// consume results
	for result := range resultChannel {
		// Print with field names for clarity
		fmt.Printf("received: value=%q err=%v\n", result)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("took time: %d ms\n", elapsed.Milliseconds())
	fmt.Printf("took time (precise): %s\n", elapsed)
}
