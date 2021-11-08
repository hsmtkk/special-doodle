package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	input := `2
Hacker
Rank
`
	want := `Hce akr
Rn ak
`
	reader := bytes.NewBufferString(input)
	var buf bytes.Buffer
	LetsReview(reader, &buf)
	got := buf.String()
	assert.Equal(t, want, got)
}
