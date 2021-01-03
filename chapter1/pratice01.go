// 커맨드라인 인수로 입력된 숫자 값들을 모두 더하는 GO 프로그램을 작성한다.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	var sum float64
	for i := 1; i < len(arguments); i++ {
		num, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum += num
		}
	}
	fmt.Println("all sum : ", sum)
}