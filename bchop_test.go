package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SortFunc func(target int, arr IntList) (int, error)

func funcTester(t *testing.T, fn SortFunc, target int, arr IntList, exceptedRet int, exceptedErr error) {
	ret, err := fn(target, arr)
	assert.Equal(t, exceptedRet, ret)
	assert.Equal(t, exceptedErr, err)
}

func innerTest(t *testing.T, fn SortFunc) {
	funcTester(t, fn, 1, []int{}, -1, nil)
	funcTester(t, fn, 1, []int{1}, 0, nil)
	funcTester(t, fn, 1, []int{2}, -1, nil)
	funcTester(t, fn, 1, []int{1, 3}, 0, nil)
	funcTester(t, fn, 3, []int{1, 3}, 1, nil)
	funcTester(t, fn, 1, []int{1, 3, 5}, 0, nil)
	funcTester(t, fn, 5, []int{1, 3, 5}, 2, nil)
	funcTester(t, fn, 1, []int{1, 3, 5, 6}, 0, nil)
	funcTester(t, fn, 3, []int{1, 3, 3, 5, 6}, 2, nil)
	funcTester(t, fn, 3, []int{1, 3, 5, 6}, 1, nil)
	funcTester(t, fn, 6, []int{1, 3, 5, 6}, 3, nil)
	funcTester(t, fn, 10, []int{1, 3, 5, 7}, -1, nil)
	funcTester(t, fn, 2, []int{1, 3, 5, 7}, -1, nil)
	funcTester(t, fn, 0, []int{1, 3, 5, 7}, -1, nil)
	funcTester(t, fn, 10, []int{1, 3, 5, 7, 9}, -1, nil)
	funcTester(t, fn, 8, []int{1, 3, 5, 7, 9}, -1, nil)
	funcTester(t, fn, 2, []int{1, 3, 5, 7, 9}, -1, nil)
	funcTester(t, fn, 0, []int{1, 3, 5, 7, 9}, -1, nil)
}

func TestBChopLoop1(t *testing.T) {
	innerTest(t, BChopLoop1)
}

func TestBChopLoop2(t *testing.T) {
	innerTest(t, BChopLoop2)
}

func TestBChopRecursive(t *testing.T) {
	innerTest(t, BChopRecursive)
}

func TestBChopSplitter(t *testing.T) {
	innerTest(t, BChopSplitter)
}

func TestBChopSplitter2(t *testing.T) {
	innerTest(t, BChopSplitter2)
}
