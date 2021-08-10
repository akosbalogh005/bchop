# bchop


## Intro
Binary search with differenet implementations in Go and Python
[See details here](http://codekata.com/kata/kata02-karate-chop/)


## Algorithms

### bchop_loop1, BChopLoop1
Binary search with simple loop.

### bchop_loop2, BChopLoop2
Binary search with simple loop, in most common way. Simplier then bchop_loop1.

### bchop_recursive, BChopRecursive
Binary search with recursive function.

### bchop_splitterfunc, BChopSplitter
Binary search with inner splitter function. The splitter function returns with the next partition
to be processed.

### bchop_splitterfunc2, BChopSplitter2
Binary search with inner splitter function. The splitter function returns with the two sides of
the array and the middle element (its position an value)


## Tests

### Python
Use pytest to run unit tests

```
pytest-3 bchop.py
```

### GO
Use go test to run unit tests
```
go test
```