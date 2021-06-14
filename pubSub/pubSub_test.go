package pubsub

import (
	"testing"
)

func assertChannelInSubs(t *testing.T, subs []chan int, ch chan int) bool {
	for _, subdCh := range subs {
		if ch == subdCh {
			return true
		}
	}
	return false
}

func TestSubscribe(t *testing.T) {
	ps := NewPubSub()

	resolutionFunc := func() {}
	solved, sln, ch := ps.Subscribe(1, &resolutionFunc)

	res := ps.Responses[1]
	if (res.solved != solved) || (res.Solution != sln) || !(assertChannelInSubs(t, res.listeners, ch)) {
		t.Fatal("subscribe did not set element")
	}
}

func TestSubscribeAndPublishSolution(t *testing.T) {
	mockSub := 1
	mockSolution := 6
	ps := NewPubSub()

	resolutionFunc := func() {
		ps.Publish(mockSub, mockSolution)
	}
	_, _, ch := ps.Subscribe(mockSub, &resolutionFunc)

	expected := <-ch
	if mockSolution != expected {
		t.Fatal("the solutions did not match")
	}
}

func TestSubscribeToAlreadyPublishedSolution(t *testing.T) {
	mockSub := 1
	mockSolution := 6
	ps := NewPubSub()
	ps.Publish(mockSub, mockSolution)
	resolutionFunc := func() {}
	solved, sln, _ := ps.Subscribe(mockSub, &resolutionFunc)
	if (solved != true) && (sln != mockSolution) {
		t.Fatal("subscription did not return already published solution")
	}
}
