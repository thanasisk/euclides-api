package main

// yeah, yeah I know technically it is a slice ...
var fibArray []uint64

func init() {
	fibArray = make([]uint64, 2)
	fibArray[0] = 0
	fibArray[1] = 1
}

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
	if n <= uint64(len(fibArray)) {
		return fibArray[n-2] + fibArray[n-1]
	} else {
		tempFib := smartFibonacci(n-2) + smartFibonacci(n-1)
		fibArray = append(fibArray, tempFib)
		return tempFib
	}
}
