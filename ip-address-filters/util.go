package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
	"time"
)

func ip2uint(ip net.IP) uint {
	return uint(binary.BigEndian.Uint32(ip))
}

func uint2ip(bit uint) net.IP {
	ip := net.IP{0, 0, 0, 0}
	binary.BigEndian.PutUint32(ip[:], uint32(bit))
	return ip
}

func progressReport(lastTime *time.Time, ops uint, format string, a ...interface{}) {
	now := time.Now()
	elapsed := now.Sub(*lastTime)
	rate := float64(ops) / elapsed.Seconds()
	nsop := float64(elapsed.Nanoseconds()) / float64(ops)
	fmt.Printf("%v %v: %v ops/s, %.2f ns/op\n", now, fmt.Sprintf(format, a), uint64(rate), nsop)
	*lastTime = now
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
