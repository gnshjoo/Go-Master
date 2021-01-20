package b

import (
	"fmt"
	"github.com/gnshjoo/GoMaster/chapter06/a"
)

func init() {
	fmt.Println("init() b")
}

func FromB() {
	fmt.Println("fromB()")
	a.FromA()
}
