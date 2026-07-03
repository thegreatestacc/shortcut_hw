package with_mutex

import "sync"

var counter int

func IncrementWithMutex() int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for gorutines := 0; gorutines < 100; gorutines++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for increment := 0; increment < 1000; increment++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return counter
}
