package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m == 0 {
		return ErrErrorsLimitExceeded
	}
	tChan := make(chan Task)
	wg := sync.WaitGroup{}

	defer wg.Wait()
	defer close(tChan)
	var e int32

	wg.Add(n)

	for w := 0; w < n; w++ {
		go func() {
			defer wg.Done()
			for t := range tChan {
				if err := t(); err != nil {
					atomic.AddInt32(&e, 1)
				}
			}
		}()
	}

	for _, t := range tasks {
		if atomic.LoadInt32(&e) >= int32(m) {
			return ErrErrorsLimitExceeded
		}
		tChan <- t
	}

	return nil
}
