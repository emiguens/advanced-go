package sync_primitives_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

// START 1 OMIT
func TestRaceCounterFixed(t *testing.T) {
	counter := new(int)
	var mu sync.Mutex // HLxxx
	var wg sync.WaitGroup

	for producerId := 0; producerId < 10; producerId++ {
		wg.Add(1)
		// spawn a new counter
		go func(producerId int) {
			mu.Lock()         // HLxxx
			defer mu.Unlock() // HLxxx
			defer wg.Done()
			*counter = *counter + 1
		}(producerId)
	}
	wg.Wait()           // wait for the N producers
	if *counter != 10 { // check spected value
		panic("counter failed.")
	}
}

// END 1 OMIT

func TestMutexPrimitive1(t *testing.T) {
	// START 2 OMIT
	type Stats struct {
		Reads  int64
		Writes int64
	}
	// END 2 OMIT
}
func TestMutexPrimitive2(t *testing.T) {
	// START 3 OMIT
	type Stats struct {
		sync.Mutex
		Reads  int64
		Writes int64
	}
	// END 3 OMIT

	// START 4 OMIT
	s := Stats{}
	s.Lock()
	defer s.Unlock()
	// END 4 OMIT
}

// START 5 OMIT
func TestRaceCounter(t *testing.T) {
	counter := new(int32) // HLxxx
	var wg sync.WaitGroup
	// multiple producers
	for producerId := 0; producerId < 10; producerId++ {
		wg.Add(1)

		// spawn a new counter
		go func(producerId int) {
			defer wg.Done()
			atomic.AddInt32(counter, 1) // HLxxx
		}(producerId)
	}
	wg.Wait()                            // wait for the N producers
	if atomic.LoadInt32(counter) != 10 { // HLxxx
		panic("counter failed.")
	}
}

// END 5 OMIT

// START 6 OMIT
func TestRaceCounterWithChannels(t *testing.T) {
	counter := 0                 // HLxxx
	increments := make(chan int) // HLxxx
	var wg sync.WaitGroup
	for producerId := 0; producerId < 10; producerId++ {
		wg.Add(1)
		// spawn a new counter
		go func(producerId int, increments chan int) {
			defer wg.Done()
			increments <- 1 // HLxxx
		}(producerId, increments)
	}
	go func() {
		wg.Wait()
		close(increments)
	}()
	for increment := range increments { // HLxxx
		counter = counter + increment // HLxxx
	} // HLxxx
	if counter != 10 {
		panic("counter failed.")
	}
}

// END 6 OMIT