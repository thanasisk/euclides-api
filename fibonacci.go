package main

// yeah, yeah it is a slice ...
var fibArray []uint64

//fibArray = {0, 1}

func naiveFibonacci(n uint64) uint64 {
	// automatic 400 if we attempt to parse negative number
	if n <= 1 {
		return n
	} else {
		return naiveFibonacci(n-2) + naiveFibonacci(n-1)
	}
}

func smartFibonacci(n uint64) uint64 {
	if n <= 1 {
		return fibArray[n]
	}
	return naiveFibonacci(n-2) + naiveFibonacci(n-1)
}
