package with_atomic

import (
	"sync"
	"sync/atomic"
)

var counter int64

func IncrementWithAtomic() int64 {
	var wg sync.WaitGroup
	atomic.StoreInt64(&counter, 0)

	for gorutines := 0; gorutines < 100; gorutines++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for increment := 0; increment < 1000; increment++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	return atomic.LoadInt64(&counter)
}
