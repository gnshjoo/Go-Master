package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("k", true, "k")
	minusO := flag.Int("0", 1, "0")
	flag.Parse()

	valueK := *minusK
	valueO := *minusO
	valueO++

	fmt.Println("-k : ", valueK)
	fmt.Println("-O : ", valueO)

}