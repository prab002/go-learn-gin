package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	This is the demonstration for the GO_routine light weight threads
	@you will lean how the multi thread works here
*/

func Worker(url string, wg *sync.WaitGroup, resultChanel chan string) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)

	fmt.Println("Url:", url)

	resultChanel <- url

}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup                // --> this should be added as worker run in other concurrent than main thread
	resultChanel := make(chan string, 3) // -> need to pass the chanel as worker will pass the data to channel and from channel we will received data to main thread

	/*
		GO Routine life cycle
		fan out --> main thread to worker/ ( light weight thread )
		fan in --> worker thread to main thread

		main --> worker --> worker to channel --> channel to main thead
	*/

	wg.Add(3)                               // --> need to pass how many thread we are running as worker
	go Worker("don.jpg", &wg, resultChanel) // --> need to pass as the refer &wg
	go Worker("don1.jpg", &wg, resultChanel)
	go Worker("don2.jpg", &wg, resultChanel)

	wg.Wait()

	close(resultChanel)
	for result := range resultChanel {
		print("received \n:", result)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("took time: %d ms\n", elapsed.Milliseconds())
	fmt.Printf("took time (precise): %s\n", elapsed)
}
