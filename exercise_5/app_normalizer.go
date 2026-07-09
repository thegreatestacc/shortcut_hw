package main

import "context"

func NormalizeApp(ctx context.Context, in <-chan AppEvent) <-chan Event {
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
					Source: "app",
					UserID: event.DeviceID,
					Action: event.Screen,
					At:     event.EventTime,
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
