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

	runtime.GC()
	fmt.Println("")
	PrintMemUsage()
	fmt.Println("[bloom] Initializing")
	bf := bloom.NewWithEstimates(numBlacklistedIPs, 0.00000001)
	PrintMemUsage()

	fmt.Println("[bloom] Initializing:Blacklisting")
	bl := &blacklist{}
	for bl.generate() {
		ip := bl.ip()
		bf.Add(ip)
	}
	PrintMemUsage()

	rng := fastrand.RNG{}

	b.ResetTimer()

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
