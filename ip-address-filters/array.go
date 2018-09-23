package main

import (
	"fmt"
	"net"
	"time"

	"github.com/valyala/fastrand"
)

type array []bool

func array_test_sequential_1(a array) {

	fmt.Println("[array] Testing sequential 1 ...")

	lastTime := time.Now()

	for i := 0; i < (1 << 32); i++ {
		bit := uint(i)
		if a[i] {
			fmt.Println("\tmatched", uint2ip(bit))
		}

		if (bit & 0xffffff) == 0xffffff {
			progressReport(&lastTime, 0xffffff, "[array-seq1] at %v", uint2ip(bit))
		}

	}

}

func array_test_sequential_2(arr array) {

	fmt.Println("[array] Testing sequential 2 ...")

	lastTime := time.Now()
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
						fmt.Println("\tmatched", ip)
					}

				}
			}
		}

		progressReport(&lastTime, 0xffffff, "[array-seq2] at %v", ip)

	}

}

func array_test_random(a array) {

	fmt.Println("[array] Testing random ...")

	lastTime := time.Now()
	rng := fastrand.RNG{}

	for i := 0; i < (1 << 28); i++ {
		bit := uint(rng.Uint32())
		if a[bit] {
			fmt.Println("\tmatched", uint2ip(bit))
		}

		if (i & 0xffffff) == 0xffffff {
			progressReport(&lastTime, 0xffffff, "[array-rand] at %v", uint2ip(bit))
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
