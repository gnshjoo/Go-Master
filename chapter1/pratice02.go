// 커맨드라인 인수로 입력된 실수 값에 대한 평균을 구하는 GO 프로그램을 작성한다.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	var avg, sum, count float64
	for i := 1; i < len(arguments); i++ {
		num, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum += num
			count += 1
		}
	}
	avg = sum / count
	fmt.Println("all avg : ", avg)
}