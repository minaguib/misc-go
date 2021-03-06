package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
	"testing"

	"github.com/mtchavez/cuckoo"
	"github.com/valyala/fastrand"
)

func BenchmarkCuckoo(b *testing.B) {

	var memBefore, memAfter uint64
	var cf *cuckoo.Filter
	rng := fastrand.RNG{}
	fmt.Println("")

	memBefore = memUsed()
	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cf = nil
			runtime.GC()
			options := []cuckoo.ConfigOption{
				cuckoo.BucketEntries(uint(4)),
				cuckoo.BucketTotal(uint(numBlacklistedIPs)),
				cuckoo.FingerprintLength(uint(4)),
				cuckoo.Kicks(uint(10)),
			}
			cf = cuckoo.New(options...)
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkCuckoo: initialize consumed", memAfter-memBefore, "bytes of memory")

	memBefore = memUsed()
	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				cf.Insert(ip)
			}
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkCuckoo: blacklist consumed", memAfter-memBefore, "bytes of memory")

	b.Run("sequential ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			binary.BigEndian.PutUint32(ip[:], uint32(i))
			if cf.Lookup(ip) {
				checkMatch(ip)
			}
		}
	})

	b.Run("random ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			bit := uint(rng.Uint32())
			binary.BigEndian.PutUint32(ip[:], uint32(bit))
			if cf.Lookup(ip) {
				checkMatch(ip)
			}
		}
	})

}
