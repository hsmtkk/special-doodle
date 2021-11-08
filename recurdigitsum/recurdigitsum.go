package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func superDigit(n string, k int32) int32 {
	ansStr, err := SuperDigit2(strings.Repeat(n, int(k)))
	if err != nil {
		log.Fatal(err)
	}
	ans, err := strconv.Atoi(ansStr)
	if err != nil {
		log.Fatal(err)
	}
	return int32(ans)
}

func SuperDigit2(n string) (string, error) {
	if len(n) == 1 {
		return n, nil
	}
	s := 0
	for _, ns := range strings.Split(n, "") {
		n, err := strconv.Atoi(ns)
		if err != nil {
			return "", err
		}
		s += n
	}
	return SuperDigit2(strconv.Itoa(s))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
