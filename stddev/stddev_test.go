package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	arr := []int32{2, 5, 2, 7, 4}
	want := float32(math.Sqrt(18.0 / 5.0))
	got := StdDev(arr)
	assert.InDelta(t, want, got, 0.1)
}

func Test1(t *testing.T) {
	arr := []int32{10, 40, 30, 50, 20}
	want := float32(14.1)
	got := StdDev(arr)
	assert.InDelta(t, want, got, 0.1)
}
