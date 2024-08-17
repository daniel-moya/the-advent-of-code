package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("Test Part 1 - Input Test", func(t *testing.T) {
		result := Run("input.txt")
		expected := 6440
		assert.Equal(t, result, expected)
	})

	t.Run("Test Part 1 - Complete Test", func(t *testing.T) {
		result := Run("input.txt")
		expected := 6440
		assert.Equal(t, result, expected)
	})

}
