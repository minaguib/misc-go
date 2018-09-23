package main

import "runtime"

func main() {

	cuckoo_test()

	runtime.GC()
	PrintMemUsage()

	hash_test()

	runtime.GC()
	PrintMemUsage()

	array_test()

	runtime.GC()
	PrintMemUsage()

	bitset_test()

	runtime.GC()
	PrintMemUsage()

	bloom_test()

	runtime.GC()
	PrintMemUsage()

}
