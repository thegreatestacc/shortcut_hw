package main

import (
	"context"
	"time"
)

func NormalizeWeb(ctx context.Context, in <-chan WebEvent) <-chan Event {
	out := make(chan Event)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case event, ok := <-in:
				if !ok {
					return
				}
				normalizedEvent := Event{
					Source: "web",
					UserID: event.SessionID,
					Action: event.URL,
					At:     time.Unix(event.TS, 0),
				}
				select {
				case <-ctx.Done():
					return
				case out <- normalizedEvent:
				}
			}
		}
	}()

	return out
}
