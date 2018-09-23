package main

import (
	"fmt"
	"time"
)

type progress struct {
	firstTime   time.Time
	maxDuration time.Duration
	ops         uint
	next        uint
}

func newProgress() *progress {
	now := time.Now()
	pro := &progress{
		firstTime:   now,
		maxDuration: (10 * time.Second),
		next:        2000000,
	}
	return pro
}

func (pro *progress) prefix(cb_or_prefix interface{}) string {
	switch x := cb_or_prefix.(type) {
	case func() string:
		return x()
	case string:
		return x
	default:
		return ""
	}
}

func (pro *progress) op() bool {
	pro.ops++
	return pro.ops >= pro.next
}

func (pro *progress) report(cb_or_prefix interface{}) bool {

	prefix := pro.prefix(cb_or_prefix)

	now := time.Now()
	elapsed := now.Sub(pro.firstTime)
	rate := float64(pro.ops) / elapsed.Seconds()
	nsop := float64(elapsed.Nanoseconds()) / float64(pro.ops)
	fmt.Printf("%v %v: %v ops/s, %.2f ns/op\n", now, prefix, uint(rate), nsop)

	pro.next = pro.ops + uint(rate)

	if elapsed < pro.maxDuration {
		return true
	} else {
		return false
	}

}
