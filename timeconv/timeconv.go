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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	return ConvertTime(s)
}

func ConvertTime(time string) string {
	hour := time[:2]
	min := time[3:5]
	sec := time[6:8]
	ampm := time[8:10]

	switch ampm {
	case "AM":
		if hour == "12" {
			return fmt.Sprintf("00:%s:%s", min, sec)
		}
		return fmt.Sprintf("%s:%s:%s", hour, min, sec)
	case "PM":
		if hour == "12" {
			return fmt.Sprintf("12:%s:%s", min, sec)
		}
		hourValue, _ := strconv.Atoi(hour)
		hourValue += 12
		return fmt.Sprintf("%02d:%s:%s", hourValue, min, sec)
	default:
		return ""
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
