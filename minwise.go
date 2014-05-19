package minhash

import (
	"hash"
	"math"
)

// MinWise is a collection of minimum hashes for a set
type MinWise struct {
	minimums []uint64
	h1       hash.Hash64
	h2       hash.Hash64
}

// NewMinWise returns a new MinWise Hashsing implementation
func NewMinWise(h1, h2 hash.Hash64, size int) *MinWise {

	minimums := make([]uint64, size)
	for i := range minimums {
		minimums[i] = math.MaxUint64
	}

	return &MinWise{
		h1:       h1,
		h2:       h2,
		minimums: minimums,
	}
}

// Push adds an element to the set.
func (m *MinWise) Push(b []byte) {

	m.h1.Reset()
	m.h1.Write(b)
	v1 := m.h1.Sum64()

	m.h2.Reset()
	m.h2.Write(b)
	v2 := m.h2.Sum64()

	for i, v := range m.minimums {
		hv := v1 + uint64(i)*v2
		if hv < v {
			m.minimums[i] = hv
		}
	}
}

// Merge combines the signatures of the second set, creating the signature of their union.
func (m *MinWise) Merge(m2 *MinWise) {

	for i, v := range m2.minimums {

		if v < m.minimums[i] {
			m.minimums[i] = v
		}
	}
}

// Cardinality estimates the cardinality of the set
func (m *MinWise) Cardinality() int {

	// http://www.cohenwang.com/edith/Papers/tcest.pdf

	sum := 0.0

	for _, v := range m.minimums {
		sum += -math.Log(float64(math.MaxUint64-v) / float64(math.MaxUint64))
	}

	return int(float64(len(m.minimums)-1) / sum)
}

// Signature returns a signature for the set.
func (m *MinWise) Signature() []uint64 {
	return m.minimums
}

// Similarity computes an estimate for the similarity between the two sets.
func (m *MinWise) Similarity(m2 *MinWise) float64 {

	if len(m.minimums) != len(m2.minimums) {
		panic("minhash minimums size mismatch")
	}

	intersect := 0

	for i := range m.minimums {
		if m.minimums[i] == m2.minimums[i] {
			intersect++
		}
	}

	return float64(intersect) / float64(len(m.minimums))
}
