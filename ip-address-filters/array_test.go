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

	runtime.GC()
	fmt.Println("")
	PrintMemUsage()
	fmt.Println("[array] Initializing")
	a := make([]bool, 1<<32)
	PrintMemUsage()

	fmt.Println("[array] Initializing:Blacklisting")
	bl := &blacklist{}
	for bl.generate() {
		ip := bl.ip()
		a[ip2uint(ip)] = true
	}
	PrintMemUsage()

	rng := fastrand.RNG{}

	b.ResetTimer()

	b.Run("sequential int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := uint(i)
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
