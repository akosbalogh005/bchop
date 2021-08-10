#!/usr/bin/python

import time
import re
import unittest
from typing import Tuple

############################################################################
def bchop_parameter_validator(function):
    """Decorator for parameter validation
       Validates incoming array and target. Additionaly handles None and empty array 
    """
    def wrapper(target, arr):
        if arr is None or target is None:
            return -1
        assert type(arr) == list
        if len(arr) == 0:
            return -1
        for i in arr:
             assert type(i) == int
        assert arr == sorted(arr)
        assert type(target) == int
        return function(target, arr)
    return wrapper




############################################################################
def _bchop_func_splitter2(arr : list) -> Tuple[list, int, int, list]:
    midpos = int((len(arr)) / 2)
    mid = arr[midpos]
    left = arr[0:midpos]
    right = arr[midpos+1:]
    return (left, midpos, mid, right)

@bchop_parameter_validator
def bchop_splitterfunc2(target : int, arr : list) -> int:
    """
    Binary search with inner splitter function. The splitter function returns with
    the two sides of the array and the middle element (its position an value)
    """
    current_arr_pos = 0
    while True:
        left, midpos, mid, right = _bchop_func_splitter2(arr)
        if mid == target:
            return current_arr_pos + midpos
        elif mid < target:
            if len(right) < 1:
                return -1
            arr = right
            current_arr_pos = current_arr_pos + midpos + 1
        elif mid > target:
            if len(left) < 1:
                return -1
            arr = left




############################################################################
def _bchop_func_splitter(current_arr : list, target) -> Tuple[list, int]:
    pmid = int((len(current_arr)) / 2)
    if current_arr[pmid] > target:
        return (current_arr[0:pmid], 0)
    return (current_arr[pmid:], pmid)

@bchop_parameter_validator
def bchop_splitterfunc(target : int, arr : list) -> int:
    """
    Binary search with inner splitter function. The splitter function returns with
    the next partition to be processed.
    """
    current_start = 0
    current_arr = arr
    while True:
        if current_arr[0] == target:
            return current_start
        elif current_arr[-1] == target:
            return current_start + len(current_arr) - 1
        if len(current_arr) <= 2:
            return -1
        current_arr, nextpos = _bchop_func_splitter(current_arr, target)
        current_start += nextpos


############################################################################
def _bchop_recursive_part(target : int, pstart : int, pend : int, arr : list ) -> int:
    if pend < pstart:
        return -1
    pmid = int((pend + pstart) / 2)
    if arr[pmid] == target:
        return pmid
    if arr[pmid] < target:
        return _bchop_recursive_part(target, pmid+1, pend, arr)   
    return _bchop_recursive_part(target, pstart, pmid-1, arr)   

@bchop_parameter_validator
def bchop_recursive(target : int, arr : list) -> int:
    """
    Binary search with recursive
    """
    return  _bchop_recursive_part(target, 0, len(arr)-1, arr)
    


############################################################################
@bchop_parameter_validator
def bchop_loop2(target : int, arr : list) -> int:
    """
    Binary search with simple loop, simplier than bchop_loop1
    """
    pstart = 0
    pend = len(arr)-1
        
    while pstart <= pend:
        pmid = int((pend + pstart) / 2)
        if arr[pmid] < target:
            pstart = pmid + 1   
        elif arr[pmid] > target:
            pend = pmid - 1
        else:
            return pmid 
    return -1



############################################################################
@bchop_parameter_validator
def bchop_loop1(target : int, arr : list) -> int:
    """
    Binary search with simple loop
    """
    pstart = 0
    pend = len(arr)-1
        
    while True:
        if arr[pstart] == target:
            return pstart
        elif arr[pend] == target:
            return pend 
        pmid = int((pend + pstart) / 2)
        if arr[pmid] <= target and pmid != pstart:
            pstart = pmid   
        elif arr[pmid] > target and pmid != pend:
            pend = pmid
        else:
            return -1 



if __name__ == '__main__':
    pass


class TestBChop(unittest.TestCase):
    def test_bchop_func2(self):
        self._test_inner(bchop_splitterfunc2)    
    
    def test_bchop_func(self):
        self._test_inner(bchop_splitterfunc)    

    def test_bchop_loop2(self):
        self._test_inner(bchop_loop2)    
    
    def test_bchop_loop1(self):
        self._test_inner(bchop_loop1)
        
    def test_bchop_recursive(self):
        self._test_inner(bchop_recursive)        
    
    def _test_inner(self, funcname):
        self.assertEqual(funcname(1, None), -1)   
        self.assertEqual(funcname(1, []), -1)     
        self.assertEqual(funcname(1, [1]), 0)
        self.assertEqual(funcname(1, [2]), -1)
        self.assertEqual(funcname(1, [1, 3]), 0)
        self.assertEqual(funcname(3, [1, 3]), 1)
        self.assertEqual(funcname(1, [1, 3, 5]), 0)
        self.assertEqual(funcname(5, [1, 3, 5]), 2)
        self.assertEqual(funcname(1, [1, 3, 5, 6]), 0)
        self.assertEqual(funcname(3, [1, 3, 5, 6]), 1)
        self.assertEqual(funcname(6, [1, 3, 5, 6]), 3)
        self.assertEqual(funcname(10, [1,3,5,7]), -1)
        self.assertEqual(funcname(2, [1,3,5,7]), -1)
        self.assertEqual(funcname(0, [1,3,5,7]), -1)
        self.assertEqual(funcname(10, [1,3,5,7,9]), -1)
        self.assertEqual(funcname(8, [1,3,5,7,9]), -1)
        self.assertEqual(funcname(2, [1,3,5,7,9]), -1)
        self.assertEqual(funcname(0, [1,3,5,7,9]), -1)