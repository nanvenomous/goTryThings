package pubsub

import (
	"sync"
)

type Response struct {
	solved    bool
	Solution  int
	listeners []chan int
}

type PubSub struct {
	Responses map[int]*Response
	mu        sync.RWMutex
	WG        sync.WaitGroup
}

func NewPubSub() *PubSub {
	ps := &PubSub{}
	ps.Responses = make(map[int]*Response)
	return ps
}

func (ps *PubSub) Subscribe(sub int, resolutionMethodP *func()) (bool, int, chan int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	res := ps.Responses[sub]

	if res == nil {
		res = &Response{false, 0, make([]chan int, 0)}
		ps.Responses[sub] = res
		ps.WG.Add(1)
		resolutionMethod := *resolutionMethodP
		go resolutionMethod()
	}

	if res.solved {
		return true, res.Solution, nil
	}

	ch := make(chan int)
	res.listeners = append(res.listeners, ch)
	return false, 0, ch
}

func (ps *PubSub) Publish(sub int, sln int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	res := ps.Responses[sub]
	if res == nil {
		res = &Response{true, sln, make([]chan int, 0)}
		ps.Responses[sub] = res
	} else {
		res.solved = true
		res.Solution = sln
	}
	for _, ch := range res.listeners {
		ch <- sln
	}
}
