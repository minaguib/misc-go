package main

import (
	"testing"
)

func BenchmarkArray(b *testing.B) {

	a := [10]int{}
	ai := 0

	for i := 0; i < b.N; i++ {
		a[ai] = i
		ai++
		if ai >= 10 {
			ai = 0
		}

	}

}
