package quartiles_test

import (
	"bytes"
	"testing"

	"github.com/hsmtkk/special-doodle/quartiles"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	input := []int{9, 5, 7, 1, 3}
	want := []int{2, 5, 8}
	got := quartiles.Quartiles(input)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	input := []int{1, 3, 5, 7}
	want := []int{2, 4, 6}
	got := quartiles.Quartiles(input)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	input := []int{3, 7, 8, 5, 12, 14, 21, 13, 18}
	want := []int{6, 12, 16}
	got := quartiles.Quartiles(input)
	assert.Equal(t, want, got)
}

func TestSplitString(t *testing.T) {
	s := "1 2 3"
	want := []int{1, 2, 3}
	got, err := quartiles.SplitString(s)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestSolve(t *testing.T) {
	input := `9
3 7 8 5 12 14 21 13 18
`
	want := `6
12
16
`
	reader := bytes.NewBufferString(input)
	var writer bytes.Buffer
	err := quartiles.Solve(reader, &writer)
	got := writer.String()
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
