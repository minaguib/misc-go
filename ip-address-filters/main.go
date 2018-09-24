package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
)


func main() {

	available := map[string]func(){
		"cuckoo": cuckoo_test,
		"bitset": bitset_test,
		"hash":   hash_test,
		"array":  array_test,
		"bloom":  bloom_test,
	}

	tests := os.Args[1:]
	if len(tests) == 0 {
		for k := range available {
			tests = append(tests, k)
		}
		sort.Strings(tests)
	}

	for _, test := range tests {
		fn := available[test]
		if fn == nil {
			panic("Unknown test:" + test)
		}
		runtime.GC()
		fmt.Println("")
		fn()
	}

}
