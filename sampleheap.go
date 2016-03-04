package sample

import (
	"math"
	"math/rand"
	"time"
)

type expSample struct {
	v Weighted
	s float64
}

// sampleHeap used for chosing k least exponential samples from sequence of
// Weighted elements. Probably deserves its own package.
type sampleHeap struct {
	h    []expSample
	rand *rand.Rand
}

func newSampleHeap(k int) *sampleHeap {
	return &sampleHeap{
		h:    make([]expSample, 0, k),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (sh *sampleHeap) Push(x Weighted) {
	esmpl := sh.rand.ExpFloat64() / x.Weight()
	if len(sh.h) < cap(sh.h) {
		sh.h = append(sh.h, expSample{v: x, s: esmpl})
		if n := len(sh.h); n == cap(sh.h) { // time to heapify
			for i := n/2 - 1; i >= 0; i-- {
				sh.down(i, n)
			}
		}
		return
	}
	if esmpl < sh.h[0].s {
		n := len(sh.h) - 1
		sh.swap(0, n)
		sh.down(0, n)
		sh.h[n] = expSample{v: x, s: esmpl}
		sh.up(n)
	}
}

func (sh *sampleHeap) Result() []Weighted {
	res := make([]Weighted, 0, len(sh.h))
	for _, s := range sh.h {
		res = append(res, s.v)
	}
	return res
}

func (sh *sampleHeap) less(i, j int) bool {
	return sh.h[i].s < sh.h[j].s || math.IsNaN(sh.h[i].s) && !math.IsNaN(sh.h[j].s)
}

func (sh *sampleHeap) swap(i, j int) {
	sh.h[i], sh.h[j] = sh.h[j], sh.h[i]
}

func (sh *sampleHeap) up(i int) {
	for {
		j := (i - 1) >> 1
		if j < 0 || sh.less(i, j) {
			return
		}
		sh.swap(i, j)
		i = j
	}
}

func (sh *sampleHeap) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			return
		}
		j := j1
		if j2 := j1 + 1; j2 < n && sh.less(j1, j2) {
			j = j2
		}
		if sh.less(j, i) {
			return
		}
		sh.swap(i, j)
		i = j
	}
}
