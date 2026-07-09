package main

import (
	"context"
	"sync"
)

func MergeChannels(ctx context.Context, inputChannels ...<-chan Event) <-chan Event {
	out := make(chan Event)

	var wg sync.WaitGroup
	wg.Add(len(inputChannels))

	for _, inputChan := range inputChannels {

		go func(in <-chan Event) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- event:
					}

				}
			}
		}(inputChan)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
