package sample

import (
	"sort"
	"testing"
)

type cnt struct {
	val   int
	count int
}

type byCount []cnt

func (c byCount) Len() int {
	return len(c)
}

func (c byCount) Less(i, j int) bool {
	return c[i].count < c[j].count
}

func (c byCount) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type byVal []cnt

func (c byVal) Len() int {
	return len(c)
}

func (c byVal) Less(i, j int) bool {
	return c[i].val < c[j].val
}

func (c byVal) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type weightedInt struct {
	elt    int
	weight float64
}

func (w weightedInt) Weight() float64 {
	return w.weight
}

func testSampleFunc(t *testing.T, sampleSize int, sf func([]Weighted, int) ([]Weighted, error)) {
	counter := make(map[int]int)
	var wl []Weighted
	for i := 1; i < 8; i++ {
		wl = append(wl, weightedInt{i, float64(i)})
	}
	for i := 0; i < 1024; i++ {
		res, err := sf(wl, sampleSize)
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range res {
			v, ok := e.(weightedInt)
			if !ok {
				t.Fatalf("type is not weightedInt: %T", e)
			}
			counter[v.elt]++
		}
	}
	var cntSlice byCount
	for v, c := range counter {
		cntSlice = append(cntSlice, cnt{val: v, count: c})
	}
	sort.Sort(cntSlice)
	if !sort.IsSorted(byVal(cntSlice)) {
		t.Fatalf("slice should be sorted by value: %v", cntSlice)
	}
}

func TestSample(t *testing.T) {
	testSampleFunc(t, 3, Sample)
}

func TestSampleChoiceFallback(t *testing.T) {
	testSampleFunc(t, 1, Sample)
}

func TestChoice(t *testing.T) {
	sf := func(w []Weighted, _ int) ([]Weighted, error) {
		c, err := Choice(w)
		return []Weighted{c}, err
	}
	testSampleFunc(t, 1, sf)
}

func BenchmarkSample(b *testing.B) {
	var wl []Weighted
	for i := 1; i < 4096; i++ {
		wl = append(wl, weightedInt{i, float64(i)})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Sample(wl, 1024); err != nil {
			b.Fatal(err)
		}
	}
}
