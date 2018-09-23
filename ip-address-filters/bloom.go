package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/valyala/fastrand"
	"github.com/willf/bloom"
)

func bloom_test_sequential(bf *bloom.BloomFilter) {

	fmt.Println("[bloom] Testing sequential ...")

	pro := newProgress()
	ip := net.IP{0, 0, 0, 0}

	for a := 0; a <= 255; a++ {
		ip[0] = uint8(a)
		for b := 0; b <= 255; b++ {
			ip[1] = uint8(b)
			for c := 0; c <= 255; c++ {
				ip[2] = uint8(c)
				for d := 0; d <= 255; d++ {
					ip[3] = uint8(d)

					if bf.Test(ip) {
						checkMatch(ip)
					}

					if pro.op() && !pro.report("[bloom-seq] at "+ip.String()) {
						return
					}

				}
			}
		}
	}

}

func bloom_test_random(bf *bloom.BloomFilter) {

	fmt.Println("[bloom] Testing random ...")

	pro := newProgress()
	rng := fastrand.RNG{}
	ip := net.IP{0, 0, 0, 0}

	for i := 0; i < (1 << 28); i++ {
		bit := uint(rng.Uint32())

		binary.BigEndian.PutUint32(ip[:], uint32(bit))
		if bf.Test(ip) {
			checkMatch(ip)
		}

		if pro.op() && !pro.report("[bloom-rand] at "+ip.String()) {
			return
		}

	}

}

func bloom_test() {

	PrintMemUsage()
	fmt.Println("[bloom] Initializing")
	bf := bloom.NewWithEstimates(256*2, 0.0001)
	PrintMemUsage()

	// Mark 256 IPs as set:
	for i := 0; i <= 255; i++ {
		ii := uint8(i)
		ip := net.IP{ii, ii, ii, ii}
		bf.Add(ip)
	}

	bloom_test_sequential(bf)
	bloom_test_random(bf)

}
