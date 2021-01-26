package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// 뮤텍스란 일종의 상호 배제용 락이다.
// 뮤텍스가 0이면 잠기지 않은 뮤텍스란 뜻이다.

// 한 번 사용한 뮤텍스는 복사하면 안 된다.

var (
	m sync.Mutex
	v1 int
)

func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1% 10 == 0 {
		v1 = v1 - 10 * i
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	fmt.Printf("%d ", read())
	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf("-> %d", read())
		}(i)
	}
	waitGroup.Wait()
	fmt.Printf("-> %d\n", read())
}