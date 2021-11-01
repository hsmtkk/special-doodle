package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	values := []int32{1, 2, 3}
	freqs := []int32{3, 2, 1}
	want := 1
	got := int(InterQuartile(values, freqs))
	assert.Equal(t, want, got)
}

func Test0(t *testing.T) {
	values := []int32{6, 12, 8, 10, 20, 16}
	freqs := []int32{5, 4, 3, 2, 1, 5}
	want := 9
	got := int(InterQuartile(values, freqs))
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	values := []int32{6, 12, 8, 10, 20, 16}
	freqs := []int32{5, 6, 7, 8, 9, 10}
	want := 8
	got := int(InterQuartile(values, freqs))
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	values := []int32{10, 40, 30, 50, 20, 10, 40, 30, 50, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 10, 40, 30, 50, 20, 10, 40, 30, 50}
	freqs := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 40, 30, 50, 20, 10, 40, 30, 50, 20}
	want := 30
	got := int(InterQuartile(values, freqs))
	assert.Equal(t, want, got)
}

func TestQuartile0(t *testing.T) {
	values := []int32{6, 6, 6, 6, 6, 8, 8, 8, 10, 10}
	want := float32(7.0)
	got := Quartile(values)
	assert.Equal(t, want, got)
}

func TestQuartile1(t *testing.T) {
	values := []int32{12, 12, 12, 12, 16, 16, 16, 16, 16, 20}
	want := float32(16.0)
	got := Quartile(values)
	assert.Equal(t, want, got)
}
