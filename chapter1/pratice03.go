// STOP이란 단어를 입력할 때까지 계속해서 입력된 정수 값을 읽는 GO 프로그램을 작성한다.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			break
		} else {
			fmt.Println(scanner.Text())
		}
	}
}