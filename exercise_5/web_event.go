package main

type WebEvent struct {
	SessionID string
	URL       string
	TS        int64 // unix seconds
}
