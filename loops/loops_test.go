package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := `2 x 1 = 2
2 x 2 = 4
2 x 3 = 6
2 x 4 = 8
2 x 5 = 10
2 x 6 = 12
2 x 7 = 14
2 x 8 = 16
2 x 9 = 18
2 x 10 = 20
`
	var buf bytes.Buffer
	PrintFirst10Multiples(2, &buf)
	got := buf.String()
	assert.Equal(t, want, got)
}
