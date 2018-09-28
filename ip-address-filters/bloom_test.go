package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
	"testing"

	"github.com/valyala/fastrand"
	"github.com/willf/bloom"
)

func BenchmarkBloom(b *testing.B) {

	var memBefore, memAfter uint64
	var bf *bloom.BloomFilter
	rng := fastrand.RNG{}
	fmt.Println("")

	memBefore = memUsed()
	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bf = nil
			runtime.GC()
			bf = bloom.NewWithEstimates(numBlacklistedIPs, 0.00000001)
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkBloom: initialize consumed", memAfter-memBefore, "bytes of memory")

	memBefore = memUsed()
	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				bf.Add(ip)
			}
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkBloom: blacklist consumed", memAfter-memBefore, "bytes of memory")

	b.Run("sequential ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			binary.BigEndian.PutUint32(ip[:], uint32(i))
			if bf.Test(ip) {
				checkMatch(ip)
			}
		}
	})

	b.Run("random ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			bit := uint(rng.Uint32())
			binary.BigEndian.PutUint32(ip[:], uint32(bit))
			if bf.Test(ip) {
				checkMatch(ip)
			}
		}
	})

}
