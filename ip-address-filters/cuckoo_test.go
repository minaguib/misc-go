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

	runtime.GC()
	fmt.Println("")
	PrintMemUsage()
	fmt.Println("[cuckoo] Initializing")
	options := []cuckoo.ConfigOption{
		cuckoo.BucketEntries(uint(2)),
		cuckoo.BucketTotal(uint(1000000)),
		cuckoo.FingerprintLength(uint(6)),
		cuckoo.Kicks(uint(500)),
	}
	cf := cuckoo.New(options...)
	PrintMemUsage()

	fmt.Println("[cuckoo] Initializing:Blacklisting")
	bl := &blacklist{}
	for bl.generate() {
		ip := bl.ip()
		cf.Insert(ip)
	}
	PrintMemUsage()

	rng := fastrand.RNG{}

	b.ResetTimer()

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
