package main

import (
	"fmt"
	"github.com/gnshjoo/GoMaster/chapter06/aPackage"

)

func main() {
	fmt.Println("Using aPackage!")
	aPackage.A()
	aPackage.B()
	fmt.Println(aPackage.MyConstant)
}