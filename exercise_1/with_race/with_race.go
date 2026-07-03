package with_race

import "sync"

var counter int

func IncrementWithRace() int {
	counter = 0
	var wg sync.WaitGroup

	for gorutines := 0; gorutines < 100; gorutines++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter += 1
			}
			// what happend if wg.Done() was here? discuss it
		}()
	}

	wg.Wait()
	return counter
}
