package routines

import (
	"sync"
	"testing"
)

func TestPrependArgument(t *testing.T) {
	messages := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go One(messages)
	wg.Add(1)
	go Two(messages)
	wg.Wait()
}
