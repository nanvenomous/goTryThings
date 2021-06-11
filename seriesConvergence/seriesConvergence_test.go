package seriesconvergence

import (
	"fmt"
	"testing"
)

func assertEqualInt(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Fatal(fmt.Sprintf("Expected: %v != Actual: %v", expected, actual))
	}
}

func TestNextElementOdd(t *testing.T) {
	expected := 10
	n := 3
	actual := NextElement(n)
	assertEqualInt(t, actual, expected)
}

func TestNextElementEven(t *testing.T) {
	expected := 2
	n := 4
	actual := NextElement(n)
	assertEqualInt(t, actual, expected)
}

// 10: 13 40 20 10 5 16 8 4 2 1
func TestChainLength(t *testing.T) {
	mem := make(map[int]int)
	expected := 10
	n := 13
	actual := ChainLength(mem, n)
	assertEqualInt(t, actual, expected)
}

func TestGetLongestChainLengthUnder_1000000(t *testing.T) {
	expected := 837799
	actual, _ := LongestChainUnder(1000000)
	assertEqualInt(t, expected, actual)
}

// // Search for a pattern in output
// func TestVariosChainLengths(t *testing.T) {
// 	mem := make(map[int]int)
// 	for n := 1; n <= 41; n += 2 {
// 		fmt.Println(ChainLength(mem, n))
// 	}
// }
