package main

import (
	"encoding/binary"
	"net"
	"testing"

	"github.com/valyala/fastrand"
	"github.com/willf/bitset"
)

func BenchmarkBitset(b *testing.B) {

	var bs *bitset.BitSet
	rng := fastrand.RNG{}

	b.Run("initialize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bs = bitset.New(1 << 32)
		}
	})

	b.Run("blacklist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bl := &blacklist{}
			for bl.generate() {
				ip := bl.ip()
				bs.Set(ip2uint(ip))
			}
		}
	})

	b.Run("sequential int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := uint(i)
			if bs.Test(bit) {
				checkMatch(uint2ip(bit))
			}
		}
	})

	b.Run("sequential ip", func(b *testing.B) {
		ip := net.IP{0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			binary.BigEndian.PutUint32(ip[:], uint32(i))
			bit := ip2uint(ip)
			if bs.Test(bit) {
				checkMatch(ip)
			}
		}
	})

	b.Run("random int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bit := uint(rng.Uint32())
			if bs.Test(bit) {
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
			if bs.Test(bit) {
				checkMatch(ip)
			}
		}
	})

}
