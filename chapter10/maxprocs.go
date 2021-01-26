package main

import (
	"fmt"
	"runtime"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPORCS: %d\n", getGOMAXPROCS())
}
