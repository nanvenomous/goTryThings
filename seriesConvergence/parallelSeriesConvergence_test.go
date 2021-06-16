package seriesconvergence

import (
	"fmt"
	"testing"
	"time"
)

// 10: 13 40 20 10 5 16 8 4 2 1
func TestChainLengthParallel(t *testing.T) {
	mem := NewMemory()
	expected := 10
	n := 13
	actual := ChainLengthParallel(mem, n)
	assertEqualInt(t, actual, expected)
}

func TestGetLongestChainLengthParallelUnder_1000000(t *testing.T) {
	expected := 837799
	now := time.Now()
	actual, _ := LongestChainParallelUnder(1000000)
	fmt.Printf("Parallel memoized solution took %f Seconds\n", time.Since(now).Seconds())
	assertEqualInt(t, expected, actual)
}
