package main

// A brute-force birthday paradox harness

import (
	"flag"
	"math/rand"
	"time"
)

var (
	classSize  = flag.Int("c", 50, "The class size (number of students)")
	numClasses = flag.Int("n", 1000, "The number of classes/iterations")
	r          = rand.New(rand.NewSource(time.Now().UnixNano()))
	from       = time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	to         = time.Date(1983, 12, 31, 23, 59, 59, 999999999, time.UTC)
	window     = to.Sub(from)
)

func assembleClass(classSize int) (numCollissionBdays int, numCollisionStudents int) {
	acc := make(map[int]int)
	for i := 0; i < classSize; i++ {
		btime := from.Add(time.Duration(r.Int63n(int64(window))))
		bday := (int(btime.Month()) << 5) | btime.Day()
		acc[bday]++
		switch acc[bday] {
		case 1:
		case 2:
			numCollissionBdays++
			numCollisionStudents += 2
		default:
			numCollisionStudents++
		}
	}
	return numCollissionBdays, numCollisionStudents
}

func main() {

	flag.Parse()

	r := make(report)
	for i := 0; i < *numClasses; i++ {
		numCollissionBdays, numCollisionStudents := assembleClass(*classSize)
		r[reportKey{numCollissionBdays, numCollisionStudents}]++
	}
	r.print()

}
