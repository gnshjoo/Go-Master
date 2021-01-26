package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 ok"
	}()

	select {
	case res :=  <-c1:
		fmt.Println(res)
	case <- time.After(time.Second *1):
		fmt.Println("timeout c1")
	}


	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "c2 ok"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("timout c2")
	}
}
