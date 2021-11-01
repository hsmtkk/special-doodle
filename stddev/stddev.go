package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'stdDev' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func stdDev(arr []int32) {
	sd := StdDev(arr)
	fmt.Printf("%.1f\n", sd)
}

func StdDev(arr []int32) float32 {
	mean := Mean(arr)
	var sum float32 = 0
	for _, a := range arr {
		sum += Pow(float32(a) - mean)
	}
	return float32(math.Sqrt(float64(sum) / float64(len(arr))))
}

func Mean(arr []int32) float32 {
	var sum int32 = 0
	for _, a := range arr {
		sum += a
	}
	return float32(sum) / float32(len(arr))
}

func Pow(f float32) float32 {
	return f * f
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	valsTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var vals []int32

	for i := 0; i < int(n); i++ {
		valsItemTemp, err := strconv.ParseInt(valsTemp[i], 10, 64)
		checkError(err)
		valsItem := int32(valsItemTemp)
		vals = append(vals, valsItem)
	}

	stdDev(vals)
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
