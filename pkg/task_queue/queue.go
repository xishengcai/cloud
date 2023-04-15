package task_queue

import "time"

type WaitQueue struct {
	metadata
}

type ProcessingQueue struct {
	metadata
}

type ResultQueue struct {
	metadata
}

type metadata struct {
	Data      interface{}
	Type      string
	UUID      string
	TimeStamp time.Time
}
