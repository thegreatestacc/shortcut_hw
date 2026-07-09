package main

import "time"

// --- Общий нормализованный вид ---

type Event struct {
	Source string
	UserID string
	Action string
	At     time.Time
}
