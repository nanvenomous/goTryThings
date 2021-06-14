package seriesconvergence

import (
	"testing"

	pubsub "github.com/mrgarelli/goTryThings/pubSub"
)

// 10: 13 40 20 10 5 16 8 4 2 1
func TestChainLengthPubSub(t *testing.T) {
	ps := pubsub.NewPubSub()
	expected := 10
	n := 13
	actual := ChainLengthPubSub(ps, n)
	assertEqualInt(t, actual, expected)
}

// Takes over 14 Seconds
// func TestPubSublGetLongestChainLengthUnder_1000000(t *testing.T) {
// 	expected := 837799
// 	now := time.Now()
// 	actual, _ := LongestChainPubSubUnder(1000000)
// 	fmt.Printf("PubSub solution took %f Seconds\n", time.Since(now).Seconds())
// 	assertEqualInt(t, expected, actual)
// }
