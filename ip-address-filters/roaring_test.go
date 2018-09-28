package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
	"testing"

	"github.com/RoaringBitmap/roaring"
	"github.com/valyala/fastrand"
)

func BenchmarkRoaring(b *testing.B) {

	var memBefore, memAfter uint64
	var r *roaring.Bitmap
	rng := fastrand.RNG{}
	fmt.Println("")

	memBefore = memUsed()
	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r = nil
			runtime.GC()
			r = roaring.New()
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkRoaring: initialize consumed", memAfter-memBefore, "bytes of memory")

	memBefore = memUsed()
	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				r.Add(ip2uint32(ip))
			}
		}
	})
	memAfter = memUsed()

	fmt.Println("BenchmarkRoaring: blacklist consumed", memAfter-memBefore, "bytes of memory")

	b.Run("sequential int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := uint32(i)
			if r.Contains(bit) {
				checkMatch(uint2ip(uint(bit)))
			}
		}
	})

	b.Run("sequential ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			binary.BigEndian.PutUint32(ip[:], uint32(i))
			bit := ip2uint32(ip)
			if r.Contains(bit) {
				checkMatch(ip)
			}
		}
	})

	b.Run("random int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := rng.Uint32()
			if r.Contains(bit) {
				checkMatch(uint2ip(uint(bit)))
			}
		}
	})

	b.Run("random ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			bit := rng.Uint32()
			binary.BigEndian.PutUint32(ip[:], bit)
			if r.Contains(bit) {
				checkMatch(ip)
			}
		}
	})

}
