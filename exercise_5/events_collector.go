package main

import "time"

func Batch(in <-chan Event, size int, timeout time.Duration) <-chan []Event {
	out := make(chan []Event)

	go func() {
		defer close(out)

		buf := make([]Event, 0, size)
		timer := time.NewTimer(timeout)

		if !timer.Stop() {
			<-timer.C
		}

		timerActive := false

		flush := func() {
			if len(buf) == 0 {
				return
			}
			b := make([]Event, len(buf))
			copy(b, buf)
			out <- b
			buf = buf[:0]
		}

		stopTimer := func() {
			if !timerActive {
				return
			}

			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			timerActive = false
		}

		resetTimer := func() {
			stopTimer()
			timer.Reset(timeout)
			timerActive = true
		}

		for {
			var timerC <-chan time.Time
			if timerActive {
				timerC = timer.C
			}
			select {
			case event, ok := <-in:
				if !ok {
					stopTimer()
					flush()
					return
				}
				if len(buf) == 0 {
					resetTimer()
				}
				buf = append(buf, event)

				if len(buf) >= size {
					stopTimer()
					flush()
				}
			case <-timerC:
				timerActive = false
				flush()
			}
		}
	}()

	return out
}
