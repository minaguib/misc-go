package main

import (
	"fmt"
	"sort"
)

type reportKey [2]int
type reportKeys []reportKey
type report map[reportKey]int

func (r report) keys() reportKeys {
	keys := make(reportKeys, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}
	return keys
}

func (r report) print() {
	fmt.Println("collision_bdays", "collision_students", "num_classes")
	for _, k := range r.sortedKeys() {
		fmt.Println(k[0], k[1], r[k])
	}
}

func (r report) sortedKeys() reportKeys {
	keys := r.keys()
	sort.Sort(keys)
	return keys
}

func (rk reportKeys) Len() int {
	return len(rk)
}

func (rk reportKeys) Less(i, j int) bool {
	if (rk[i][0] < rk[j][0]) || (rk[i][0] == rk[j][0] && rk[i][1] < rk[j][1]) {
		return true
	}
	return false
}

func (rk reportKeys) Swap(i, j int) {
	rk[i], rk[j] = rk[j], rk[i]
}
