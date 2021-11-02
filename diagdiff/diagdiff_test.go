package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	want := int32(15)
	got := diagonalDifference(arr)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	arr := [][]int32{
		{-1, 1, -7, -8},
		{-10, -8, -5, -2},
		{0, 9, 7, -1},
		{4, 4, -2, 1},
	}
	want := int32(1)
	got := diagonalDifference(arr)
	assert.Equal(t, want, got)
}
