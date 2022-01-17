//Other than main go routine, launch two go routines and use WaitGroup to synchronize them
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go printHelloWorld()
	go printHelloIndia()
	wg.Wait()
}

func printHelloWorld() {
	defer wg.Done()
	fmt.Println("Hello World")
	for i := 0; i < 100; i++ {
		fmt.Println(i)

	}
}

func printHelloIndia() {
	defer wg.Done()
	fmt.Println("Hello India")
	for i := 0; i < 100; i++ {
		fmt.Println(i)

	}
}
