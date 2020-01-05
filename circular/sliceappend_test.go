package main

import (
	"testing"
)

func BenchmarkSliceAppend(b *testing.B) {

	s := make([]int, 10)

	for i := 0; i < b.N; i++ {
		s = s[1:]
		s = append(s, i)
	}

}
