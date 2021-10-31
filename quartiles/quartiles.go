package quartiles

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

	got := Quartiles(ns)
	q1 := got[0]
	q2 := got[1]
	q3 := got[2]

	bw := bufio.NewWriter(output)
	defer bw.Flush()

	bw.WriteString(fmt.Sprintf("%d\n", q1))
	bw.WriteString(fmt.Sprintf("%d\n", q2))
	bw.WriteString(fmt.Sprintf("%d\n", q3))

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

func Quartiles(array []int) []int {
	sort.Ints(array)
	center := len(array) / 2
	var lower []int
	var upper []int
	var q1 int
	var q2 int
	var q3 int
	if len(array)%2 == 0 {
		lower = array[:center]
		upper = array[center:]
		q2 = (array[center-1] + array[center]) / 2
	} else {
		lower = array[:center]
		upper = array[center+1:]
		q2 = array[center]
	}
	q1 = Median(lower)
	q3 = Median(upper)
	return []int{q1, q2, q3}
}

func Median(array []int) int {
	center := len(array) / 2
	if len(array)%2 == 0 {
		return (array[center-1] + array[center]) / 2
	} else {
		return array[center]
	}
}
