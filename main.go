package main

import (
	"fmt"
	"time"
)

/*
	This is the demonstration for the GO_routine light weight threads
	@you will lean how the multi thread works here
*/

func Worker(url string) string {
	time.Sleep(50 * time.Millisecond)

	fmt.Println("Url:", url)
	return url
}

func main() {
	startTime := time.Now()

	result := Worker("don.jpg")
	fmt.Println("this is result:", result)

	elapsed := time.Since(startTime)
	fmt.Printf("took time: %d ms\n", elapsed.Milliseconds())
	fmt.Printf("took time (precise): %s\n", elapsed)
}
