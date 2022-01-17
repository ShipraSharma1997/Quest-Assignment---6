//Write a program thatLaunches 10 go routines, Each go routine adds 10 numbers to the channel, Pull the numbers off the channel and print them
package main

import (
	"fmt"
)

func main() {
	maxGoroutines := 10
	guard := make(chan struct{}, maxGoroutines)

	for i := 0; i < 20; i++ {
		guard <- struct{}{} // would block if guard channel is already filled
		go func(n int) {
			worker(n)
			<-guard
		}(i)
	}
}

func worker(i int) { fmt.Println("doing work on", i) }
