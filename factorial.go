package main

import "math"

func naiveFactorial(n uint64) uint64 {
	if n > 0 {
		res := n * naiveFactorial(n-1)
		return res
	}
	return 1
}

// bet you were expecting the same memoization trick
// nope! Golang's STDlib FTW
// https://en.wikipedia.org/wiki/Gamma_function
func smartFactorial(n uint64) uint64 {
	// values are capped
	return uint64(math.Gamma(float64(n + 1)))
}
