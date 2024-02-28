package day1

import "testing"

func TestRun(t *testing.T) {
	t.Run("#1 Example", func(t *testing.T) {
		handleTest("input-1-test.txt", t, 142)
	})

	t.Run("#1 Answer", func(t *testing.T) {
		handleTest("input-1.txt", t, 54390)
	})
}

func handleTest(path string, t *testing.T, expectedTotal int) {
	var total = Run(path)
	if total != expectedTotal {
		t.Errorf("AOC#2 Assert posisble games cube sum, Expected %v , Found %v", expectedTotal, total)
	}
}
