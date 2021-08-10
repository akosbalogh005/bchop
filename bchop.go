package main

import (
	"fmt"
	"sort"
)

type IntList []int

// checkParam general Param checking
func checkParam(target int, arr IntList) error {
	if !sort.IntsAreSorted(arr) {
		return fmt.Errorf("array must be sorted")
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////
//BChopRecursive Binary search with recursive
func BChopRecursive(target int, arr IntList) (int, error) {
	return bchopRecursivePart(target, 0, len(arr)-1, arr)
}

func bchopRecursivePart(target int, pstart int, pend int, arr IntList) (int, error) {
	if pend < pstart {
		return -1, nil
	}
	pmid := int((pend + pstart) / 2)
	if arr[pmid] == target {
		return pmid, nil
	}
	if arr[pmid] < target {
		return bchopRecursivePart(target, pmid+1, pend, arr)
	}
	return bchopRecursivePart(target, pstart, pmid-1, arr)
}

/////////////////////////////////////////////////////////////////////////////
//BChopLoop2 Binary search with simple loop, simplier than BChopLoop1
func BChopLoop2(target int, arr IntList) (int, error) {
	err := checkParam(target, arr)
	if err != nil {
		return -1, err
	}
	if arr == nil || len(arr) == 0 {
		return -1, nil
	}

	pstart := 0
	pend := len(arr) - 1
	for pstart <= pend {
		pmid := int((pend + pstart) / 2)
		if arr[pmid] < target {
			pstart = pmid + 1
		} else if arr[pmid] > target {
			pend = pmid - 1
		} else {
			return pmid, nil
		}
	}
	return -1, nil
}

/////////////////////////////////////////////////////////////////////////////
// BChopLoop1 Binary search with simple loop
func BChopLoop1(target int, arr IntList) (int, error) {
	err := checkParam(target, arr)
	if err != nil {
		return -1, err
	}
	if arr == nil || len(arr) == 0 {
		return -1, nil
	}

	pstart := 0
	pend := len(arr) - 1

	for {
		if arr[pstart] == target {
			return pstart, nil
		} else if arr[pend] == target {
			return pend, nil
		}
		pmid := int((pend + pstart) / 2)
		if arr[pmid] <= target && pmid != pstart {
			pstart = pmid
		} else if arr[pmid] > target && pmid != pend {
			pend = pmid
		} else {
			return -1, nil
		}
	}
}

/////////////////////////////////////////////////////////////////////////////
//BChopSplitter Binary search with inner splitter function. The splitter function returns with
//the next partition to be processed.
func BChopSplitter(target int, arr IntList) (int, error) {
	err := checkParam(target, arr)
	if err != nil {
		return -1, err
	}
	if arr == nil || len(arr) == 0 {
		return -1, nil
	}
	currentStart := 0
	currentArr := arr
	for {
		if currentArr[0] == target {
			return currentStart, nil
		} else if currentArr[len(currentArr)-1] == target {
			return currentStart + len(currentArr) - 1, nil
		}
		if len(currentArr) <= 2 {
			return -1, nil
		}
		var nextpos int
		currentArr, nextpos = bChopSplitterFunc(currentArr, target)
		currentStart += nextpos
	}
}

func bChopSplitterFunc(currentArr IntList, target int) (IntList, int) {

	pmid := int((len(currentArr)) / 2)
	if currentArr[pmid] > target {
		return currentArr[0:pmid], 0
	}
	return currentArr[pmid:], pmid
}

/////////////////////////////////////////////////////////////////////////////
//BChopSplitter2 Binary search with inner splitter function. The splitter function returns with
// the two sides of the array and the middle element (its position an value)
func BChopSplitter2(target int, arr IntList) (int, error) {
	err := checkParam(target, arr)
	if err != nil {
		return -1, err
	}
	if arr == nil || len(arr) == 0 {
		return -1, nil
	}
	current_arr_pos := 0
	for {
		left, midpos, mid, right := bChopSplitterFunc2(arr)
		if mid == target {
			return current_arr_pos + midpos, nil
		} else if mid < target {
			if len(right) < 1 {
				return -1, nil
			}
			arr = right
			current_arr_pos = current_arr_pos + midpos + 1
		} else if mid > target {
			if len(left) < 1 {
				return -1, nil
			}
			arr = left
		}
	}
}

func bChopSplitterFunc2(arr IntList) (IntList, int, int, IntList) {
	midpos := int((len(arr)) / 2)
	mid := arr[midpos]
	left := arr[0:midpos]
	right := arr[midpos+1:]
	return left, midpos, mid, right

}

func main() {
	ret, err := BChopSplitter2(6, []int{1, 3, 5, 6})
	fmt.Printf("Result: %v, error: %v", ret, err)
}
