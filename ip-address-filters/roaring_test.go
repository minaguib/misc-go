package main

import (
	"encoding/binary"
	"net"
	"runtime"
	"testing"

	"github.com/RoaringBitmap/roaring"
	"github.com/valyala/fastrand"
)

func BenchmarkRoaring(b *testing.B) {

	var r *roaring.Bitmap
	rng := fastrand.RNG{}

	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r = nil
			runtime.GC()
			r = roaring.New()
		}
	})

	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				r.Add(ip2uint32(ip))
			}
		}
	})

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
