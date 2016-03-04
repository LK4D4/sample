// Package sample solves problem of choosing k random elements with weights
// from some(potentially big) sequence. Such problem is called Reservoir Sampling.
package sample

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	// ErrEmptyInput used when input sequence is empty - nothing to sample.
	ErrEmptyInput = errors.New("input sequence empty")
	// ErrTooBigSample returned when requested sample is bigger than input sequence.
	ErrTooBigSample = errors.New("requested sample size exceeds size of sequence")
)

// Weighted is interface for "weighted" elements in sequence.
type Weighted interface {
	Weight() float64
}

// Sample samples random k elements with respect to their weights from slice of
// Weighted elements.
//
// Complexity is O(n*log(k)), O(k) memory is used.
func Sample(input []Weighted, k int) ([]Weighted, error) {
	if len(input) == 0 {
		return nil, ErrEmptyInput
	}
	if k >= len(input) {
		return nil, ErrTooBigSample
	}
	if k == 1 {
		ch, err := Choice(input)
		if err != nil {
			return nil, err
		}
		return []Weighted{ch}, nil
	}
	sh := newSampleHeap(k)
	for _, w := range input {
		sh.Push(w)
	}
	return sh.Result(), nil
}

// Choice returns Weighted element with respect to its weight from slice of
// Weighted elements.
//
// Complexity is O(n), memory usage is O(1).
func Choice(input []Weighted) (Weighted, error) {
	if len(input) == 0 {
		return nil, ErrEmptyInput
	}
	var wsum float64
	for _, w := range input {
		wsum += w.Weight()
	}
	threshold := rand.Float64() * wsum
	for _, w := range input {
		threshold -= w.Weight()
		if threshold <= 0 {
			return w, nil
		}
	}
	return nil, fmt.Errorf("internal error")
}
