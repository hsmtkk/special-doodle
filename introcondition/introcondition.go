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

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	msg := GetMessage(N)
	fmt.Println(msg)
}

const WERID = "Weird"
const NOT_WERID = "Not Weird"

func GetMessage(n int32) string {
	if n%2 == 1 {
		return WERID
	}
	if 2 <= n && n <= 5 {
		return NOT_WERID
	}
	if 6 <= n && n <= 20 {
		return WERID
	}
	return NOT_WERID
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
