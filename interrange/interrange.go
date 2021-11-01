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
 * Complete the 'interQuartile' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY values
 *  2. INTEGER_ARRAY freqs
 */

func interQuartile(values []int32, freqs []int32) {
	// Print your answer to 1 decimal place within this function
	f := InterQuartile(values, freqs)
	fmt.Printf("%.1f\n", f)
}

func InterQuartile(values []int32, freqs []int32) float32 {
	array := []int32{}
	for i := 0; i < len(values); i++ {
		var j int32
		for j = 0; j < freqs[i]; j++ {
			array = append(array, values[i])
		}
	}
	var lower []int32
	var upper []int32
	center := len(array) / 2
	if len(array)%2 == 0 {
		lower = array[:center]
		upper = array[center:]
	} else {
		lower = array[:center-1]
		upper = array[center:]
	}
	q1 := Quartile(lower)
	q3 := Quartile(upper)
	return q3 - q1
}

func Quartile(values []int32) float32 {
	center := len(values) / 2
	if len(values)%2 == 0 {
		return float32((values[center-1] + values[center])) / 2.0
	} else {
		return float32(values[center])
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	valTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var val []int32

	for i := 0; i < int(n); i++ {
		valItemTemp, err := strconv.ParseInt(valTemp[i], 10, 64)
		checkError(err)
		valItem := int32(valItemTemp)
		val = append(val, valItem)
	}

	freqTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var freq []int32

	for i := 0; i < int(n); i++ {
		freqItemTemp, err := strconv.ParseInt(freqTemp[i], 10, 64)
		checkError(err)
		freqItem := int32(freqItemTemp)
		freq = append(freq, freqItem)
	}

	interQuartile(val, freq)
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
