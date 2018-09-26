package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
)

const numBlacklistedIPs = 1000000

func ip2uint(ip net.IP) uint {
	return uint(binary.BigEndian.Uint32(ip))
}

func uint2ip(bit uint) net.IP {
	ip := net.IP{0, 0, 0, 0}
	binary.BigEndian.PutUint32(ip[:], uint32(bit))
	return ip
}

func checkMatch(ip net.IP) {
	if isBlacklisted(ip) {
		// Correct match
	} else {
		// False-positive match
		fmt.Println(ip, "false-positive match")
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
