package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/seiflotfy/cuckoofilter"
	"github.com/valyala/fastrand"
)

func cuckoo_test_sequential(cf *cuckoofilter.CuckooFilter) {

	fmt.Println("[cuckoo] Testing sequential ...")

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

					if cf.Lookup(ip) {
						fmt.Println("\tmatched", ip)
					}

				}
			}
		}

		progressReport(&lastTime, 0xffffff, "[cuckoo-seq] at %v", ip)

	}

}

func cuckoo_test_random(cf *cuckoofilter.CuckooFilter) {

	fmt.Println("[cuckoo] Testing random ...")

	lastTime := time.Now()
	rng := fastrand.RNG{}
	ip := net.IP{0, 0, 0, 0}

	for i := 0; i < (1 << 28); i++ {
		bit := uint(rng.Uint32())
		binary.BigEndian.PutUint32(ip[:], uint32(bit))
		if cf.Lookup(ip) {
			fmt.Println("\tmatched", ip)
		}

		if (i & 0xffffff) == 0xffffff {
			progressReport(&lastTime, 0xffffff, "[cuckoo-rand] at %v", uint2ip(bit))
		}
	}

}

func cuckoo_test() {

	PrintMemUsage()
	fmt.Println("[cuckoo] Initializing")
	cf := cuckoofilter.NewCuckooFilter(1 << 25)
	PrintMemUsage()

	// Mark 256 IPs as set:
	for i := 0; i <= 255; i++ {
		ii := uint8(i)
		ip := net.IP{ii, ii, ii, ii}
		cf.Insert(ip)
	}

	cuckoo_test_sequential(cf)
	cuckoo_test_random(cf)

}
