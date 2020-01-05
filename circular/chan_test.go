package main

import (
	"testing"
)

func BenchmarkChan(b *testing.B) {

	c := make(chan int, 10)

	for i := 0; i < b.N; i++ {
		c <- i
		<-c
	}

}
