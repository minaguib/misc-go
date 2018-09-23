package main

import (
	"fmt"
	"runtime"
)

func main() {

	bitset_test()

	runtime.GC()
	fmt.Println("")

	cuckoo_test()

	runtime.GC()
	fmt.Println("")

	hash_test()

	runtime.GC()
	fmt.Println("")

	array_test()

	runtime.GC()
	fmt.Println("")

	bloom_test()

	runtime.GC()
	fmt.Println("")

}
