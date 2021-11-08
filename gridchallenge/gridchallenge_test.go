package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	input := []string{"abc", "ade", "efg"}
	want := "YES"
	got := gridChallenge(input)
	assert.Equal(t, want, got)
}

func TestB(t *testing.T) {
	input := []string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}
	want := "YES"
	got := gridChallenge(input)
	assert.Equal(t, want, got)
}
