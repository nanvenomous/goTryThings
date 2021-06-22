package main

import (
	"fmt"
	"time"

	seriesconvergence "github.com/mrgarelli/goTryThings/seriesConvergence"
)

func getTimeUnder(maxToTest int, underTest func(int) (int, int)) float64 {
	now := time.Now()
	underTest(maxToTest)
	return time.Since(now).Seconds()
}

func main() {
	maxToTest := 1000000
	// pubSubTime := getTimeUnder(maxToTest, seriesconvergence.LongestChainPubSubUnder)
	// fmt.Printf("PubSub Time: %f Seconds\n", pubSubTime)
	parallelTime := getTimeUnder(maxToTest, seriesconvergence.LongestChainParallelUnder)
	fmt.Printf("Parallel Time: %f Seconds\n", parallelTime)
	memoizedTime := getTimeUnder(maxToTest, seriesconvergence.LongestChainUnder)
	fmt.Printf("Memoized Time Time: %f Seconds\n", memoizedTime)
}
