package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	x := []int32{1, 2, 3}
	w := []int32{5, 6, 7}
	want := 2
	got := int(WeightedMean(x, w))
	assert.Equal(t, want, got)
}
