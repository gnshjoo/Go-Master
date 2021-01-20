package aPackage

import "fmt"

func A() {
	fmt.Println("Thi is Function A!")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}

const MyConstant = 123
const privateConstant = 21


