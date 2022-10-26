package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m == 0 {
		return ErrErrorsLimitExceeded
	}

	outerInc := 1
	outerMax := 1
	if n < len(tasks) {
		outerInc = n
		outerMax = len(tasks)
	}

	errorsCount := 0
	var mux sync.Mutex
	for i := 0; i < outerMax; i += outerInc {
		var wg sync.WaitGroup
		for k := 0; k < n; k++ {
			wg.Add(1)
			go func(i, k int, wg *sync.WaitGroup, errorsCount *int) {
				defer wg.Done()
				fn := tasks[i+k]
				if fn != nil {
					err := fn()
					if err != nil {
						mux.Lock()
						*errorsCount++
						mux.Unlock()
					}
				}
			}(i, k, &wg, &errorsCount)
		}
		wg.Wait()
		if errorsCount >= m {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}
