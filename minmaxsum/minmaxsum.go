package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	s := MiniMaxSum(arr)
	fmt.Printf("%d %d\n", s.Min, s.Max)
}

type SumMinMax struct {
	Min int64
	Max int64
}

func MiniMaxSum(arr []int32) SumMinMax {
	var min int64 = 0
	var max int64 = 0
	for skip := 0; skip < 5; skip++ {
		var s int64 = 0
		for i := 0; i < 5; i++ {
			if i == skip {
				continue
			}
			s += int64(arr[i])
		}
		if min == 0 {
			min = s
		}
		if max == 0 {
			max = s
		}
		if s < min {
			min = s
		}
		if s > max {
			max = s
		}
	}
	return SumMinMax{Min: min, Max: max}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
