package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	count, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	kvs, err := ParseKeyValues(lines[1 : count+1])
	if err != nil {
		log.Fatal(err)
	}
	Lookup(lines[count+1:], kvs)
}

func ParseKeyValues(lines []string) (map[string]string, error) {
	kvs := map[string]string{}
	for _, line := range lines {
		elems := strings.Split(line, " ")
		if len(elems) != 2 {
			return nil, fmt.Errorf("invalid format: %s", line)
		}
		k := elems[0]
		v := elems[1]
		kvs[k] = v
	}
	return kvs, nil
}

func Lookup(lines []string, kvs map[string]string) {
	for _, line := range lines {
		v, ok := kvs[line]
		if ok {
			fmt.Printf("%s=%s\n", line, v)
		} else {
			fmt.Println("Not found")
		}
	}
}
