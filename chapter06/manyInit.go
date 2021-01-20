package main

import (
	"fmt"
	"github.com/gnshjoo/GoMaster/chapter06/a"
	"github.com/gnshjoo/GoMaster/chapter06/b"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	a.FromA()
	b.FromB()
}
