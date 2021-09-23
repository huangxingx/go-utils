package go_utils

import (
	"sync"
	"time"
)

// Future async object
type Future struct {
	isFinished bool
	result     interface{}
	resultChan chan interface{}
	l          sync.Mutex
}

// GetResult get return value
func (f *Future) GetResult() interface{} {
	f.l.Lock()
	defer f.l.Unlock()
	if f.isFinished {
		return f.result
	}

	select {
	// timeout
	case <-time.Tick(time.Second * 6):
		f.isFinished = true
		f.result = nil
		return nil
	case f.result = <-f.resultChan:
		f.isFinished = true
		return f.result
	}
}

// SetResult set return value
func (f *Future) SetResult(result interface{}) {
	if f.isFinished == true {
		return
	}
	f.resultChan <- result
	close(f.resultChan)
}

// NewFuture init Future
func NewFuture() *Future {
	return &Future{
		isFinished: false,
		result:     nil,
		resultChan: make(chan interface{}, 1),
	}
}
