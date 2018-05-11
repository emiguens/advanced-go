package channels_test

import "testing"

func TestBlockingChannel(t *testing.T) {

	// START 1 OMIT
	select {
	case ch <- 42: // HLxxx
		// non-blocking write
	default:
		// habria bloqueao
	}
	// END 1 OMIT

	// START 2 OMIT
	select {
	case value := <-ch: // HLxxx
		// non-blocking write
	default:
		// habria bloqueao
	}
	// END 2 OMIT
}
