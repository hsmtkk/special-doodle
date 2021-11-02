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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	rt := PlusMinus(arr)
	fmt.Printf("%.5f\n", rt.PositiveRatio)
	fmt.Printf("%.5f\n", rt.NegativeRatio)
	fmt.Printf("%.5f\n", rt.ZeroRatio)
}

type Ratio struct {
	PositiveRatio float32
	NegativeRatio float32
	ZeroRatio     float32
}

func PlusMinus(arr []int32) Ratio {
	positiveCount := 0
	negativeCount := 0
	zeroCount := 0
	for _, n := range arr {
		if n == 0 {
			zeroCount += 1
		} else if n > 0 {
			positiveCount += 1
		} else {
			negativeCount += 1
		}
	}
	arrLength := float32(len(arr))
	posRatio := float32(positiveCount) / arrLength
	negRatio := float32(negativeCount) / arrLength
	zeroRatio := float32(zeroCount) / arrLength
	return Ratio{PositiveRatio: posRatio, NegativeRatio: negRatio, ZeroRatio: zeroRatio}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
