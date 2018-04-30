package rk

import (
	"sort"
)

// RK represents a Random Key
// RK is "immutable"
type RK []float64

// Clone returns a copy
func (r RK) Clone() RK {
	clone := make(RK, len(r))
	copy(clone, r)
	return clone
}

// Sort returns a sorted clone
func (r RK) Sort() RK {
	clone := r.Clone()
	sort.Float64s(clone)
	return clone
}

// Permutation returns a []int representation of the RK
func (r RK) Permutation() []int {
	permutation := []int{}
	sorted := r.Sort()
	for _, k := range r {
		permutation = append(permutation, sort.SearchFloat64s(sorted, k))
	}
	return permutation
}

// Uniformize returns a uniformized copy
func (r RK) Uniformize() RK {
	rk := RK{}

	return rk
}

// Mean returns a new RK resulting of the mean of RKs passed in
func Mean(rks []RK) RK {
	rksLen := len(rks)
	rkLen := len(rks[0])
	result := RK{}
	for i := 0; i < rkLen; i++ {
		m := float64(0)
		for _, rk := range rks {
			m = m + rk[i]
		}
		m = m / float64(rksLen)
		result = append(result, m)
	}
	return result
}
