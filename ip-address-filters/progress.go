package main

import (
	"fmt"
	"time"
)

// Helper for reporting progress on ops within the testing code
// Kinda like benchmarking, but invoked by code instead of invokes the code
type progress struct {
	// Config
	reportEvery    time.Duration
	maxDuration    time.Duration
	firstReportOps uint

	// Internal state:
	startTime     time.Time
	ops           uint
	nextReportOps uint
}

func newProgress() *progress {
	p := &progress{
		reportEvery:    (1 * time.Second),
		maxDuration:    (10 * time.Second),
		firstReportOps: 1000000,
		startTime:      time.Now(),
	}
	p.nextReportOps = p.firstReportOps
	return p
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

// Increment number of ops
// Returns true if you should call .report() now-ish, false if not
// Intentionally kept small to allow inlining
func (pro *progress) op() bool {
	pro.ops++
	return pro.ops >= pro.nextReportOps
}

// Write to stdout a progress report
// cb_or_prefix can be a string, or a function that returns a string
// Use the function method if the prefix is computationally expensive to compute and you don't feel
// strongly about Go's short-circuiting in a case like so:
//
// pro.op() && pro.report("at client " + client.expensiveMethod())
//
// Returns a hint to keep working (true), or to stop working (false) as configured maxDuration has been reached
//
func (pro *progress) report(cb_or_prefix interface{}) bool {

	prefix := pro.prefix(cb_or_prefix)

	now := time.Now()
	elapsed := now.Sub(pro.startTime)
	opsPerSec := float64(pro.ops) / elapsed.Seconds()
	nsPerOp := float64(elapsed.Nanoseconds()) / float64(pro.ops)

	fmt.Printf("%v %v: %v ops/s, %.2f ns/op\n", now, prefix, uint(opsPerSec), nsPerOp)

	pro.nextReportOps = pro.ops + (uint(opsPerSec) * uint(pro.reportEvery.Seconds()))

	if elapsed < pro.maxDuration {
		// Hint to keep going
		return true
	} else {
		// Hint to stop - maxDuration is reached
		return false
	}

}
