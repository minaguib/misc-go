package main

import (
	"container/ring"
	"testing"
)

func BenchmarkRing(b *testing.B) {

	r := ring.New(10)

	for i := 0; i < b.N; i++ {
		r = r.Next()
		r.Value = i
	}

}
