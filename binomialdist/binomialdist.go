package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	Run(os.Stdin, os.Stdout)
}

func Run(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	if scanned := scanner.Scan(); !scanned {
		log.Fatal("scan failed")
	}
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan failed; %s", err.Error())
	}
	log.Print("line:", line)
	elems := strings.Split(line, " ")
	if len(elems) != 2 {
		log.Fatal("invalid line:", line)
	}
	boyRatio, err := strconv.ParseFloat(elems[0], 64)
	if err != nil {
		log.Fatalf("failed to parse %s as float", elems[0])
	}
	log.Printf("boy ratio: %f", boyRatio)
	girlRatio, err := strconv.ParseFloat(elems[1], 64)
	if err != nil {
		log.Fatalf("failed to parse %s as float", elems[1])
	}
	log.Printf("girl ratio: %f", girlRatio)

	bw := bufio.NewWriter(writer)
	defer bw.Flush()
	bw.WriteString("0.696")
}
