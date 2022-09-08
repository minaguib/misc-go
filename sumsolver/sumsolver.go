package main

/* Given an input set of integers and a target sum value
   Lists all permutations of set's elements (repetition allowed) that sum up to the target sum value
*/

import "fmt"
import "strings"

var SET = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
var TARGETSUM = 100

type state []int

// Utility function that converts state to a printable representation
func (s state) String() string {
	var b strings.Builder
	first := true
	for i, count := range s {
		v := SET[i]
		for x := 0; x < count; x++ {
			if !first {
				b.WriteString("+")
			}
			fmt.Fprintf(&b, "%d", v)
			first = false
		}
	}
	return b.String()
}

func permute(s state, sum int, pos int) {

	if pos >= len(s) {
		// Leaf
		if sum == TARGETSUM {
			fmt.Println(s)
		}
		return
	}

	for c := 0; ; c++ {
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
	s := make(state, len(SET))
	permute(s, 0, 0)
}
