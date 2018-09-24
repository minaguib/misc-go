package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/seiflotfy/cuckoofilter"
	"github.com/valyala/fastrand"
)

func cuckoo_test_sequential(cf *cuckoofilter.CuckooFilter) {

	fmt.Println("[cuckoo] Testing sequential ...")

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

					if cf.Lookup(ip) {
						checkMatch(ip)
					}

					if pro.op() && !pro.report("[cuckoo-seq] at "+ip.String()) {
						return
					}

				}
			}
		}

	}

}

func cuckoo_test_random(cf *cuckoofilter.CuckooFilter) {

	fmt.Println("[cuckoo] Testing random ...")

	pro := newProgress()
	rng := fastrand.RNG{}
	ip := net.IP{0, 0, 0, 0}

	for i := 0; i < (1 << 28); i++ {
		bit := uint(rng.Uint32())
		binary.BigEndian.PutUint32(ip[:], uint32(bit))
		if cf.Lookup(ip) {
			checkMatch(ip)
		}

		if pro.op() && !pro.report("[cuckoo-rand] at "+ip.String()) {
			return
		}

	}

}

func cuckoo_test() {

	PrintMemUsage()
	fmt.Println("[cuckoo] Initializing")
	cf := cuckoofilter.NewCuckooFilter(1 << 25)
	PrintMemUsage()

	fmt.Println("[cuckoo] Initializing:Blacklisting")
	bl := &blacklist{}
	for bl.generate() {
		ip := bl.ip()
		cf.Insert(ip)
	}
	PrintMemUsage()

	cuckoo_test_sequential(cf)
	cuckoo_test_random(cf)

}
