package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqual(t *testing.T) {

	tests := []struct {
		InputA bool
		InputB bool
		Want   bool
	}{
		{
			InputA: true,
			InputB: true,
			Want:   true,
		},
		{
			InputA: true,
			InputB: false,
			Want:   false,
		},
		{
			InputA: false,
			InputB: true,
			Want:   false,
		},
		{
			InputA: false,
			InputB: false,
			Want:   true,
		},
	}

	for _, test := range tests {

		assert.Equal(t, test.Want, IsEqual(test.InputA, test.InputB))
	}
}
