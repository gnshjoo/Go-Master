package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("You ar using", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using go Version", runtime.Version())
	fmt.Println("Number of CPUs : ", runtime.NumCPU())
	fmt.Println("Number of Goroutines :", runtime.NumGoroutine())

}
