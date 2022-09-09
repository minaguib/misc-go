package main

/* Given an input sorted set of integers and a target sum value
   Lists all permutations of set's elements that sum up to the target sum value
*/

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
)

type state []int

var SET = []int{}
var TARGETSUM = 0
var ALLOWREPETITION = false
var OUT = bufio.NewWriter(os.Stdout)

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

	/* Config */
	set := flag.String("set", "1,2,3,4,5,6", "the set of integers")
	sum := flag.Int("sum", 10, "the sum to aim for")
	allowrepetition := flag.Bool("allowrepetition", false, "allow repetition of single element from set")
	flag.Parse()
	ALLOWREPETITION = *allowrepetition
	TARGETSUM = *sum
	for _, v := range strings.Split(*set, ",") {
		v, _ := strconv.Atoi(v)
		SET = append(SET, v)
	}

	/* Start */
	s := make(state, len(SET))
	permute(s, 0, 0)
}
