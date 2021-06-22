package seriesconvergence

import (
	"sync"
)

type Memory struct {
	Solutions map[int]int
	mu        sync.RWMutex
}

func NewMemory() *Memory {
	mem := &Memory{}
	mem.Solutions = make(map[int]int)
	return mem
}

func (mem *Memory) Get(key int) (int, bool) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	val, keyExists := mem.Solutions[key]
	return val, keyExists
}

func (mem *Memory) Set(key int, val int) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	mem.Solutions[key] = val
}

func ChainLengthParallel(mem *Memory, n int) int {
	if n == 1 {
		return 1
	}
	if val, inMap := mem.Get(n); inMap {
		return val
	} else {
		nxtElem := NextElement(n)
		subLen := ChainLengthParallel(mem, nxtElem)
		len := subLen + 1
		mem.Set(n, len)
		return len
	}
}

func LongestChainParallelUnder(num int) (int, int) {
	mem := NewMemory()
	wg := sync.WaitGroup{}
	sem := make(chan int, 2)

	singleChain := func(i int) {
		sem <- 1
		ChainLengthParallel(mem, i)
		<-sem
		wg.Done()
	}
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go singleChain(i)
	}
	wg.Wait()
	maxKey, maxElement := KeyOfMaxValueInMap(mem.Solutions)
	return maxKey, maxElement
}
