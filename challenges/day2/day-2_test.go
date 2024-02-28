package day2

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		handleTest("input-2-test.txt", t, 8, 2286)
	})

	t.Run("Example", func(t *testing.T) {
		handleTest("input-2.txt", t, 2771, 70924)
	})
}

func handleTest(path string, t *testing.T, expectedTotal int64, expectedMinPowerProduct float64) {
	var total, minPowerProduct = Run(path)
	if total != expectedTotal {
		t.Errorf("AOC#2 Assert posisble games cube sum, Expected %v , Found %v", expectedTotal, total)
	}

	if minPowerProduct != expectedMinPowerProduct {
		t.Errorf("AOC#2 Assert min power product, Expected %v , Found %v", expectedMinPowerProduct, minPowerProduct)
	}

}
