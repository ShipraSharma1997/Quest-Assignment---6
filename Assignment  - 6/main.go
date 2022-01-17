// Golang program to find the number of logical processors used by current process
package main

import (
	"fmt"
	"runtime"
)

func main() {

	// Using the NumCPU function
	fmt.Println(runtime.NumCPU())
}
