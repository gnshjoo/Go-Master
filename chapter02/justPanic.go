package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Not enough arguments!")
	}

	fmt.Println("Tanks for the arguments(s)!")
}
