package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := int32(15)
	got := CalculateTotalCost(12.00, 20, 8)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := int32(13)
	got := CalculateTotalCost(10.25, 17, 5)
	assert.Equal(t, want, got)
}
