package flow

import (
	"sync"
	"time"
)

type Bucket struct {
	Capacity     int
	DripInterval time.Duration
	PerDrip      int
	consumed     int
	started      bool
	kill         chan bool
	m            sync.Mutex
}

type Buckets []*Bucket

var Pool []Buckets

func New(b ...*Bucket) int {
	Buckets(b).Start()
	defer Buckets(b).Stop()
	Pool = append(Pool, b)
	return len(Pool) - 1
}
