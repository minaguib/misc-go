package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
	"testing"

	"github.com/valyala/fastrand"
)

func BenchmarkHash(b *testing.B) {

	var memBefore, memAfter uint64
	var h map[uint]bool
	rng := fastrand.RNG{}
	fmt.Println("")

	memBefore = memUsed()
	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h = nil
			runtime.GC()
			h = make(map[uint]bool, numBlacklistedIPs)
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkHash: initialize consumed", memAfter-memBefore, "bytes of memory")

	memBefore = memUsed()
	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				h[ip2uint(ip)] = true
			}
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkHash: blacklist consumed", memAfter-memBefore, "bytes of memory")

	b.Run("sequential int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := uint(uint32(i))
			if h[bit] {
				checkMatch(uint2ip(bit))
			}
		}
	})

	b.Run("sequential ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			binary.BigEndian.PutUint32(ip[:], uint32(i))
			bit := ip2uint(ip)
			if h[bit] {
				checkMatch(ip)
			}
		}
	})

	b.Run("random int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := uint(rng.Uint32())
			if h[bit] {
				checkMatch(uint2ip(bit))
			}
		}
	})

	b.Run("random ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			bit := uint(rng.Uint32())
			binary.BigEndian.PutUint32(ip[:], uint32(bit))
			bit = ip2uint(ip)
			if h[bit] {
				checkMatch(ip)
			}
		}
	})

}
