package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	LetsReview(os.Stdin, os.Stdout)
}

func LetsReview(reader io.Reader, writer io.Writer) {
	lines := []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	bw := bufio.NewWriter(writer)
	defer bw.Flush()
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		bw.WriteString(ShuffleLine(line))
		bw.WriteString("\n")
	}
}

func ShuffleLine(line string) string {
	first := ""
	second := ""
	for i := 0; i < len(line); i++ {
		if i%2 == 0 {
			first += string(line[i])
		} else {
			second += string(line[i])
		}
	}
	return first + " " + second
}
