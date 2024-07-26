package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	cases := []struct {
		name     string
		val      string
		expected bool
	}{
		{"true", "001", true},
		{"false", "0a1", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := IsNumeric(c.val)
			assert.Equal(t, c.expected, actual)
		})
	}
}
