package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	arr := []int32{1, 3, 5, 7, 9}
	want := SumMinMax{Min: 16, Max: 24}
	got := MiniMaxSum(arr)
	assert.Equal(t, want, got)
}

func TestB(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5}
	want := SumMinMax{Min: 10, Max: 14}
	got := MiniMaxSum(arr)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	want := SumMinMax{Min: 2063136757, Max: 2744467344}
	got := MiniMaxSum(arr)
	assert.Equal(t, want, got)
}
