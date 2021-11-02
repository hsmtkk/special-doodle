package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	assert.Equal(t, "12:01:00", ConvertTime("12:01:00PM"))
}

func TestB(t *testing.T) {
	assert.Equal(t, "00:01:00", ConvertTime("12:01:00AM"))
}

func TestC(t *testing.T) {
	assert.Equal(t, "19:05:45", ConvertTime("07:05:45PM"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "00:40:22", ConvertTime("12:40:22AM"))
}
