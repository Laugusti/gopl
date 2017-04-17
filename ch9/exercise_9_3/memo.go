// Extend the Func tpe and the (*Memo).Get method so that callers may
// provide an optional done channel through which they can cancel the
// operation. The results of a cancelled Func call should not be cached.
package memo

import (
	"errors"
	"sync"
)

var CancelledError = errors.New("cancelled operation")

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

type entry struct {
	res       result
	cancelled bool
	ready     chan struct{} // closed when res is ready
}

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string, cancel <-chan struct{}) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil || e.cancelled {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		go func() {
			e.res.value, e.res.err = memo.f(key, cancel)
			close(e.ready)
		}()

		select {
		case <-cancel:
			e.cancelled = true
			return nil, CancelledError
		case <-e.ready:
		}

	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		select {
		case <-e.ready:
			// retry if previous entry was cancelled
			if e.cancelled {
				return memo.Get(key, cancel)
			}
		case <-cancel:
			return nil, CancelledError
		}

	}
	return e.res.value, e.res.err
}
