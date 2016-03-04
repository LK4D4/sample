# sample

[![Build Status](https://travis-ci.org/LK4D4/sample.svg?branch=master)](https://travis-ci.org/LK4D4/sample)
[![GoDoc](https://godoc.org/github.com/LK4D4/sample?status.svg)](https://godoc.org/github.com/LK4D4/sample)

Sampling with weights

# Usage

```go
package main

import (
	"fmt"

	"github.com/LK4D4/sample"
)

type weightedInt struct {
	elt    int
	weight float64
}

func (w weightedInt) Weight() float64 {
	return w.weight
}

func SampleInt(w []sample.Weighted, k int) ([]int, error) {
	s, err := sample.Sample(w, k)
	if err != nil {
		return nil, err
	}
	var res []int
	for _, wi := range s {
		res = append(res, wi.(weightedInt).elt)
	}
	return res, nil
}

func main() {
	var wl []sample.Weighted
	for i := 1; i < 32; i++ {
		wl = append(wl, weightedInt{i, float64(i)})
	}
	n, err := SampleInt(wl, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
```
