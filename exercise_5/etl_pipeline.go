// etl_pipeline.go
//
// Эталонное решение для лайв-кодинга по concurrency в Go: ETL-pipeline с fan-in.
//
//   genWeb ─┐                                         ┌─ batch(size + timeout) ─ consume
//           ├─ normalize ─┐                           │
//           │             ├─ merge (fan-in) ─ Event ──┤
//   genApp ─┘─ normalize ─┘                           └─
//
// Запуск:  go run -race etl_pipeline.go
// Тесты:   go test -race ./...
//
// Сборка pipeline вынесена в RunPipeline: она принимает каналы-источники,
// поэтому в тестах в неё можно подавать детерминированные данные, а main
// остаётся тонким entrypoint'ом, который создаёт генераторы и консьюмер.

package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// --- Сборка pipeline (нормализация -> merge -> батчинг) ---
// Принимает уже готовые каналы-источники, что делает её тестируемой.
func RunPipeline(
	ctx context.Context,
	web <-chan WebEvent,
	app <-chan AppEvent,
	batchSize int,
	batchTimeout time.Duration,
) <-chan []Event {
	// нормализация
	normalizedWebEvents := NormalizeWeb(ctx, web)
	normalizedAppEvents := NormalizeApp(ctx, app)

	// merge
	mergedEvents := MergeChannels(ctx, normalizedWebEvents, normalizedAppEvents)

	// батчинг
	return Batch(mergedEvents, batchSize, batchTimeout)
}

// --- Генераторы ---

func genWeb(ctx context.Context, every time.Duration) <-chan WebEvent {
	out := make(chan WebEvent)
	go func() {
		defer close(out)
		t := time.NewTicker(every)
		defer t.Stop()
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				i++
				ev := WebEvent{
					SessionID: "web-" + strconv.Itoa(i),
					URL:       "/p/" + strconv.Itoa(rand.Intn(5)),
					TS:        time.Now().Unix(),
				}
				// Отправку тоже прикрываем ctx, иначе на отмене
				// горутина зависнет на out <- ev, если читателя уже нет.
				select {
				case out <- ev:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func genApp(ctx context.Context, every time.Duration) <-chan AppEvent {
	out := make(chan AppEvent)
	go func() {
		defer close(out)
		t := time.NewTicker(every)
		defer t.Stop()
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				i++
				ev := AppEvent{
					DeviceID:  "dev-" + strconv.Itoa(i),
					Screen:    "screen_" + strconv.Itoa(rand.Intn(5)),
					EventTime: time.Now(),
				}
				select {
				case out <- ev:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

// --- Консьюмер ---

func consume(in <-chan []Event) {
	// TODO what is "n"? counter? maybe batchCounter is better name?
	n := 0
	for b := range in {
		n++
		web, app := 0, 0
		for _, e := range b {
			switch e.Source {
			case "web":
				web++
			case "app":
				app++
			}
		}
		log.Printf("batch #%d: %d событий (web=%d app=%d)", n, len(b), web, app)
	}
	log.Printf("готово, всего батчей: %d", n)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	web := genWeb(ctx, 80*time.Millisecond)
	app := genApp(ctx, 120*time.Millisecond)

	batches := RunPipeline(ctx, web, app, 10, 500*time.Millisecond)

	consume(batches) // блокирует main, пока pipeline не сольётся до конца
}

//func normalizeWeb(ctx context.Context, ch <-chan WebEvent) <-chan Event{
//	return NormalizeWeb(ctx, ch)
//}
//
//func normalizeApp(ctx context.Context, ch <-chan AppEvent) <-chan Event {
//	return NormalizeApp(ctx, ch)
//}
