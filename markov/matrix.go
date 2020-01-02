package markov

import (
	"math/rand"
)

// Boundary is the value that indicates boundary between sentences (usually represented visually by period and/or newline(s))
var Boundary = ""

type row struct {
	cols  map[string]uint
	total uint
}

// Matrix represents the Markov matrix
type Matrix struct {
	rows     map[string]*row
	lastFrom string
}

// NewMatrix initializes and returns a new Matrix
func NewMatrix() *Matrix {
	return &Matrix{
		rows:     make(map[string]*row),
		lastFrom: Boundary,
	}
}

// AddPair records in the Markov matrix the probability increase of transitioning from `from` to `to`
func (m *Matrix) AddPair(from string, to string) {

	r, ok := m.rows[from]
	if !ok {
		r = &row{
			cols: make(map[string]uint),
		}
		m.rows[from] = r
	}

	r.cols[to]++
	r.total++
}

// Add is a helper function that calls AddPair with the supplied `to` and the previous call's `to` as the `from`
func (m *Matrix) Add(to string) {
	m.AddPair(m.lastFrom, to)
	m.lastFrom = to
}

// Get picks from the matrix a random `to` following the supplied `from`
func (m *Matrix) Get(from string) string {
	r, ok := m.rows[from]
	if !ok {
		//fmt.Printf("[%v] not found\n")
		return Boundary
	}
	ra := uint(rand.Intn(int(r.total))) + 1
	for to, weight := range r.cols {
		//fmt.Printf("[%v]/[%v] -> [%v]/[%v] (%v)\n", from, r.total, to, weight, ra)
		if ra <= weight {
			return to
		}
		ra -= weight
	}
	return Boundary
}
