package main

import (
	"encoding/binary"
	"net"
)

const delta = (1 << 32) / numBlacklistedIPs

type blacklist struct {
	last uint32
}

func (b *blacklist) generate() bool {
	if ((1 << 32) - 1 - delta) > b.last {
		b.last += delta
		return true
	}
	return false
}

func (b *blacklist) ip() net.IP {
	ip := net.IP{0, 0, 0, 0}
	binary.BigEndian.PutUint32(ip[:], b.last)
	return ip
}

func isBlacklisted(ip net.IP) bool {
	u := ip2uint(ip)
	return (u % delta) == 0
}
