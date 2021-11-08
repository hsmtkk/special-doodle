package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	orig := "middle-Outz"
	want := "okffng-Qwvb"
	got := caesarCipher(orig, 2)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	orig := "www.abc.xy"
	want := "fff.jkl.gh"
	got := caesarCipher(orig, 87)
	assert.Equal(t, want, got)
}
