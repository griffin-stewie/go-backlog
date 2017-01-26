// Generated by: gen
// TypeWriter: slice
// Directive: +gen on *Version

package gobacklog

import (
	"errors"
	"math/rand"
)

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// VersionSlice is a slice of type *Version. Use it where you would use []*Version.
type VersionSlice []*Version

// Where returns a new VersionSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv VersionSlice) Where(fn func(*Version) bool) (result VersionSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Count gives the number elements of VersionSlice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv VersionSlice) Count(fn func(*Version) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// SortBy returns a new ordered VersionSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv VersionSlice) SortBy(less func(*Version, *Version) bool) VersionSlice {
	result := make(VersionSlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortVersionSlice(result, less, 0, n, maxDepth)
	return result
}

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv VersionSlice) GroupByString(fn func(*Version) string) map[string]VersionSlice {
	result := make(map[string]VersionSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// GroupByInt groups elements into a map keyed by int. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv VersionSlice) GroupByInt(fn func(*Version) int) map[int]VersionSlice {
	result := make(map[int]VersionSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// GroupByBool groups elements into a map keyed by bool. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv VersionSlice) GroupByBool(fn func(*Version) bool) map[bool]VersionSlice {
	result := make(map[bool]VersionSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv VersionSlice) First(fn func(*Version) bool) (result *Version, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no VersionSlice elements return true for passed func")
	return
}

// MaxBy returns an element of VersionSlice containing the maximum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
func (rcv VersionSlice) MaxBy(less func(*Version, *Version) bool) (result *Version, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the MaxBy of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if rcv[i] != rcv[m] && !less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// MinBy returns an element of VersionSlice containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
func (rcv VersionSlice) MinBy(less func(*Version, *Version) bool) (result *Version, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Min of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// Distinct returns a new VersionSlice whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv VersionSlice) Distinct() (result VersionSlice) {
	appended := make(map[*Version]bool)
	for _, v := range rcv {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// DistinctBy returns a new VersionSlice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv VersionSlice) DistinctBy(equal func(*Version, *Version) bool) (result VersionSlice) {
Outer:
	for _, v := range rcv {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// Shuffle returns a shuffled copy of VersionSlice, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (rcv VersionSlice) Shuffle() VersionSlice {
	numItems := len(rcv)
	result := make(VersionSlice, numItems)
	copy(result, rcv)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[r], result[i] = result[i], result[r]
	}
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapVersionSlice(rcv VersionSlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortVersionSlice(rcv VersionSlice, less func(*Version, *Version) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapVersionSlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownVersionSlice(rcv VersionSlice, less func(*Version, *Version) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapVersionSlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortVersionSlice(rcv VersionSlice, less func(*Version, *Version) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownVersionSlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapVersionSlice(rcv, first, first+i)
		siftDownVersionSlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeVersionSlice(rcv VersionSlice, less func(*Version, *Version) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapVersionSlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapVersionSlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapVersionSlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeVersionSlice(rcv VersionSlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapVersionSlice(rcv, a+i, b+i)
	}
}

func doPivotVersionSlice(rcv VersionSlice, less func(*Version, *Version) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeVersionSlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeVersionSlice(rcv, less, m, m-s, m+s)
		medianOfThreeVersionSlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeVersionSlice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapVersionSlice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapVersionSlice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapVersionSlice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeVersionSlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeVersionSlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortVersionSlice(rcv VersionSlice, less func(*Version, *Version) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortVersionSlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotVersionSlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortVersionSlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortVersionSlice(rcv, mhi, b)
		} else {
			quickSortVersionSlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortVersionSlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortVersionSlice(rcv, less, a, b)
	}
}