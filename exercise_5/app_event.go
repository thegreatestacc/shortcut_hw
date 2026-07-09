package main

import "time"

type AppEvent struct {
	DeviceID  string
	Screen    string
	EventTime time.Time
}
