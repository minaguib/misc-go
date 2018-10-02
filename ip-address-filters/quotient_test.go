package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
	"testing"

	"github.com/Nomon/qf-go"
	"github.com/valyala/fastrand"
)

func BenchmarkQuotient(b *testing.B) {

	var memBefore, memAfter uint64
	var q *qf.QuotientFilter
	rng := fastrand.RNG{}
	fmt.Println("")

	memBefore = memUsed()
	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			q = nil
			runtime.GC()
			q = qf.NewProbability(numBlacklistedIPs, 0.00000001)
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkQuotient: initialize consumed", memAfter-memBefore, "bytes of memory")

	memBefore = memUsed()
	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				ipstr := string(ip)
				q.Add(ipstr)
			}
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkQuotient: blacklist consumed", memAfter-memBefore, "bytes of memory")

	b.Run("sequential ip string", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			binary.BigEndian.PutUint32(ip[:], uint32(i))
			ipstr := string(ip)
			if q.Contains(ipstr) {
				checkMatch(ip)
			}
		}
	})

	b.Run("random ip string", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			bit := uint(rng.Uint32())
			binary.BigEndian.PutUint32(ip[:], uint32(bit))
			ipstr := string(ip)
			if q.Contains(ipstr) {
				checkMatch(ip)
			}
		}
	})

}
