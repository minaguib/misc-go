package main

/* Given an input sorted set of integers and a target sum value
   Lists all permutations of set's elements that sum up to the target sum value
*/

import (
	"bufio"
	"os"
	"strconv"
)

var SET = []int{1, 2, 3, 4, 5, 6}
var TARGETSUM = 10
var ALLOWREPETITION = false
var OUT = bufio.NewWriter(os.Stdout)

type state []int

// Utility function that prints a human-friendly version of state
func (s state) print() {
	first := true
	for i, count := range s {
		v := strconv.Itoa(SET[i])
		for x := 0; x < count; x++ {
			if !first {
				OUT.WriteString("+")
			}
			OUT.WriteString(v)
			first = false
		}
	}
	OUT.WriteString("\n")
}

func permute(s state, sum int, pos int) {

	if pos >= len(s) {
		// Leaf
		if sum == TARGETSUM {
			s.print()
		}
		return
	}

	for c := 0; ALLOWREPETITION || c <= 1; c++ {
		extra := c * SET[pos]
		newsum := sum + extra
		if newsum > TARGETSUM {
			break
		}
		s[pos] = c
		permute(s, newsum, pos+1)
	}
}

func main() {
	defer OUT.Flush()
	s := make(state, len(SET))
	permute(s, 0, 0)
}
