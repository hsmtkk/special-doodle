package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	arr := []int32{0, 1, 2, 4, 6, 5, 3}
	want := int32(3)
	got := findMedian(arr)
	assert.Equal(t, want, got)
}
