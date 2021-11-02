package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := int32(1)
	got := lonelyinteger([]int32{1})
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := int32(2)
	got := lonelyinteger([]int32{1, 1, 2})
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	want := int32(2)
	got := lonelyinteger([]int32{0, 0, 1, 2, 1})
	assert.Equal(t, want, got)
}
