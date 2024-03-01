package day3

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		handleTest("input-3-test.txt", t, 4361)
	})
}

func handleTest(
	path string,
	t *testing.T,
	expectedTotal int,
) {
	var total = Run(path)
	if total != expectedTotal {
		t.Errorf("AOC#2 Assert part numbers sum, Expected %v , Found %v", expectedTotal, total)
	}
}
