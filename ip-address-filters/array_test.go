package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
	"testing"

	"github.com/valyala/fastrand"
)

func BenchmarkArray(b *testing.B) {

	var a []bool
	var memBefore, memAfter uint64
	rng := fastrand.RNG{}
	fmt.Println("")

	memBefore = memUsed()
	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a = nil
			runtime.GC()
			a = make([]bool, 1<<32)
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkArray: initialize consumed", memAfter-memBefore, "bytes of memory")

	memBefore = memUsed()
	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				a[ip2uint(ip)] = true
			}
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkArray: blacklist consumed", memAfter-memBefore, "bytes of memory")

	b.Run("sequential int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := uint(uint32(i))
			if a[bit] {
				checkMatch(uint2ip(bit))
			}
		}
	})

	b.Run("sequential ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			binary.BigEndian.PutUint32(ip[:], uint32(i))
			bit := ip2uint(ip)
			if a[bit] {
				checkMatch(ip)
			}
		}
	})

	b.Run("random int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := uint(rng.Uint32())
			if a[bit] {
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
			if a[bit] {
				checkMatch(ip)
			}
		}
	})

}
