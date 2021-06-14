package seriesconvergence

import (
	pubsub "github.com/mrgarelli/goTryThings/pubSub"
)

func ChainLengthPubSub(ps *pubsub.PubSub, n int) int {
	if n == 1 {
		return 1
	}

	onFirstSubscription := func() {
		nxtElem := NextElement(n)
		subLen := ChainLengthPubSub(ps, nxtElem)
		ps.Publish(n, subLen+1)
		ps.WG.Done()
	}
	solved, sln, ch := ps.Subscribe(n, &onFirstSubscription)
	if solved {
		return sln
	} else {
		sln = <-ch
		return sln
	}
}

func MaxKeyValueInResponses(ps *pubsub.PubSub) (int, int) {
	maxKey := -1
	maxElement := -1
	for key, res := range ps.Responses {
		if res.Solution > maxElement {
			maxKey = key
			maxElement = res.Solution
		}
	}
	return maxKey, maxElement
}

func LongestChainPubSubUnder(num int) (int, int) {
	ps := pubsub.NewPubSub()
	singleChainCall := func(i int) {
		ChainLengthPubSub(ps, i)
		ps.WG.Done()
	}
	for i := 1; i <= num; i++ {
		ps.WG.Add(1)
		go singleChainCall(i)
	}
	ps.WG.Wait()
	maxKey, maxElement := MaxKeyValueInResponses(ps)
	return maxKey, maxElement
}
