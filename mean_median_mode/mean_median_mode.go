package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if err := Solve(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func Solve(input io.Reader, output io.Writer) error {
	lines := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read input; %w", err)
	}

	nsStr := lines[1]
	ns, err := SplitString(nsStr)
	if err != nil {
		return err
	}

	mean := Mean(ns)
	median := Median(ns)
	mode := Mode(ns)

	bw := bufio.NewWriter(output)
	defer bw.Flush()

	bw.WriteString(fmt.Sprintf("%.1f\n", mean))
	bw.WriteString(fmt.Sprintf("%.1f\n", median))
	bw.WriteString(fmt.Sprintf("%d\n", mode))

	return nil
}

func SplitString(s string) ([]int, error) {
	ns := []int{}
	for _, s := range strings.Split(s, " ") {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to parse as int; %s; %w", s, err)
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func Mean(ns []int) float64 {
	sum := 0
	for _, n := range ns {
		sum += n
	}
	return float64(sum) / float64(len(ns))
}

func Median(ns []int) float64 {
	sort.Ints(ns)
	center := len(ns) / 2
	if len(ns)%2 == 0 {
		return float64(ns[center-1]+ns[center]) / 2.0
	} else {
		return float64(ns[center])
	}
}

func Mode(ns []int) int {
	counter := map[int]int{}
	for _, n := range ns {
		count, ok := counter[n]
		if ok {
			counter[n] = count + 1
		} else {
			counter[n] = 1
		}
	}
	maxCount := 1
	for _, count := range counter {
		if count > maxCount {
			maxCount = count
		}
	}
	modes := []int{}
	for n, count := range counter {
		if count == maxCount {
			modes = append(modes, n)
		}
	}
	sort.Ints(modes)
	return modes[0]
}
