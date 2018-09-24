package main

import (
	"fmt"
	"net"

	"github.com/valyala/fastrand"
)

type hash map[uint]bool

func hash_test_sequential_1(h hash) {

	fmt.Println("[hash] Testing sequential 1 ...")

	pro := newProgress()

	for i := 0; i < (1 << 32); i++ {
		bit := uint(i)

		if h[bit] {
			checkMatch(uint2ip(bit))
		}

		if pro.op() && !pro.report("[hash-seq1] at "+uint2ip(bit).String()) {
			return
		}

	}

}

func hash_test_sequential_2(h hash) {

	fmt.Println("[hash] Testing sequential 2 ...")

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
					bit := ip2uint(ip)

					if h[bit] {
						checkMatch(ip)
					}

					if pro.op() && !pro.report("[hash-seq2] at "+ip.String()) {
						return
					}

				}
			}
		}
	}

}

func hash_test_random(h hash) {

	fmt.Println("[hash] Testing random ...")

	pro := newProgress()
	rng := fastrand.RNG{}

	for i := 0; i < (1 << 28); i++ {
		bit := uint(rng.Uint32())

		if h[bit] {
			checkMatch(uint2ip(bit))
		}

		if pro.op() && !pro.report("[hash-rand] at "+uint2ip(bit).String()) {
			return
		}

	}

}

func hash_test() {

	PrintMemUsage()
	fmt.Println("[hash] Initializing")
	h := make(hash, numBlacklistedIPs)

	PrintMemUsage()
	fmt.Println("[hash] Initializing:Blacklisting")
	bl := &blacklist{}
	for bl.generate() {
		ip := bl.ip()
		h[ip2uint(ip)] = true
	}
	PrintMemUsage()

	hash_test_sequential_1(h)
	hash_test_sequential_2(h)
	hash_test_random(h)

}
