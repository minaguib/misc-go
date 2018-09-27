package main

import (
	"encoding/binary"
	"net"
	"runtime"
	"testing"

	"github.com/valyala/fastrand"
	"github.com/willf/bloom"
)

func BenchmarkBloom(b *testing.B) {

	var bf *bloom.BloomFilter
	rng := fastrand.RNG{}

	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bf = nil
			runtime.GC()
			bf = bloom.NewWithEstimates(numBlacklistedIPs, 0.00000001)
		}
	})

	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				bf.Add(ip)
			}
		}
	})

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
