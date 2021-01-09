package main

import "C"
import "fmt"

//#include <stdio.h>
//void callC(){
//printf("calling C code!\n");
//}

func main() {
	fmt.Println("A go statement!")
	C.callC()
	fmt.Println("Another Go statment!")
}

// 생각처럼 작동이 안된다 . 알아봐야함
