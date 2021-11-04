package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := []int{0, 1}
	got := twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := []int{1, 2}
	got := twoSum([]int{3, 2, 4}, 6)
	assert.Equal(t, want, got)
}
