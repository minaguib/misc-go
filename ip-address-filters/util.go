package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func ip2uint32(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip)
}

func ip2uint(ip net.IP) uint {
	return uint(ip2uint32(ip))
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
