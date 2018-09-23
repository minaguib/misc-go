package main

import (
	"fmt"
	"net"

	"github.com/valyala/fastrand"
)

type array []bool

func array_test_sequential_1(a array) {

	fmt.Println("[array] Testing sequential 1 ...")

	pro := newProgress()

	for i := 0; i < (1 << 32); i++ {

		bit := uint(i)
		if a[i] {
			checkMatch(uint2ip(bit))
		}

		if pro.op() && !pro.report("[array-seq1] at "+uint2ip(bit).String()) {
			return
		}

	}

}

func array_test_sequential_2(arr array) {

	fmt.Println("[array] Testing sequential 2 ...")

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
					if arr[bit] {
						checkMatch(ip)
					}

					if pro.op() && !pro.report("[array-seq2] at "+ip.String()) {
						return
					}

				}
			}
		}
	}

}

func array_test_random(a array) {

	fmt.Println("[array] Testing random ...")

	pro := newProgress()
	rng := fastrand.RNG{}

	for i := 0; i < (1 << 28); i++ {

		bit := uint(rng.Uint32())
		if a[bit] {
			checkMatch(uint2ip(bit))
		}

		if pro.op() && !pro.report("[array-rand] at "+uint2ip(bit).String()) {
			return
		}

	}

}

func array_test() {

	PrintMemUsage()
	fmt.Println("[array] Initializing")
	a := make(array, 1<<32)
	PrintMemUsage()

	// Mark 256 IPs as set:
	for i := 0; i <= 255; i++ {
		ii := uint8(i)
		ip := net.IP{ii, ii, ii, ii}
		a[ip2uint(ip)] = true
	}

	array_test_sequential_1(a)
	array_test_sequential_2(a)
	array_test_random(a)

}
