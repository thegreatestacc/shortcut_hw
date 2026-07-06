package default_woker_pool

import (
	"fmt"
	"sync"
)

// thread-safe statistic struct
type Statistics struct {
	mu                 sync.Mutex
	panicsCounter      int
	failedTasksCounter int
	// workerID to successful tasks
	processed map[int]int
}

func NewStatistics() *Statistics {
	return &Statistics{
		processed: make(map[int]int),
	}
}

func (s *Statistics) IncrementPanicsCount() {
	s.mu.Lock()
	s.panicsCounter += 1
	s.mu.Unlock()
}

func (s *Statistics) IncrementFailedTasksCounter() {
	s.mu.Lock()
	s.failedTasksCounter += 1
	s.mu.Unlock()
}

func (s *Statistics) IncrementProcessed(workerID int) {
	s.mu.Lock()
	s.processed[workerID] += 1
	s.mu.Unlock()
}

func (s *Statistics) Print() {
	s.mu.Lock()
	fmt.Printf(
		"Statistic: panics %d\n failed tasks %d\n processed: %v\n",
		s.panicsCounter, s.failedTasksCounter, s.processed)
	s.mu.Unlock()
}
