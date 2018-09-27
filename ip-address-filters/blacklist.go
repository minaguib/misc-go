package main

import (
	"encoding/binary"
	"net"
)

const numBlacklistedIPs = 42949672 // 1% of the ipv4 space - 2^32 * 0.01
const numClusters = 1000000        // Across 1M networked clusters

const numPerCluster = numBlacklistedIPs / numClusters
const clustersOffset = (1 << 32) / numClusters

type blacklist struct {
	cluster uint
	seq     uint
	started bool
}

func (b *blacklist) generate() bool {

	if !b.started {
		// First
		b.started = true
		return true
	}

	if b.seq < (numPerCluster - 1) {
		b.seq++
		return true
	} else if b.cluster < (numClusters - 1) {
		b.seq = 0
		b.cluster++
		return true
	}

	return false
}

func (b *blacklist) ip() net.IP {
	ip := net.IP{0, 0, 0, 0}
	u := (b.cluster * clustersOffset) + b.seq
	binary.BigEndian.PutUint32(ip[:], uint32(u))
	return ip
}

func isBlacklisted(ip net.IP) bool {

	u := ip2uint(ip)

	// Round down to nearest multiple of clustersOffset
	o1 := (u / clustersOffset) * clustersOffset
	// Then figure out the last blacklisted ip in cluster:
	o2 := o1 + numPerCluster - 1

	return u >= o1 && u <= o2

}
