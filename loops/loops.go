package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	PrintFirst10Multiples(n, os.Stdout)
}

func PrintFirst10Multiples(n int32, writer io.Writer) {
	bw := bufio.NewWriter(writer)
	defer bw.Flush()
	var i int32 = 1
	for i = 1; i <= 10; i++ {
		s := fmt.Sprintf("%d x %d = %d\n", n, i, n*i)
		bw.WriteString(s)
	}
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
