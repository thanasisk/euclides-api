package main

var factorials []uint64

func init() {
	factorials = make([]uint64, 2)
	factorials[0] = 1
	factorials[1] = 1
}

func naiveFactorial(n uint64) uint64 {
	if n > 0 {
		res := n * naiveFactorial(n-1)
		return res
	}
	return 1
}

// FIXME
func smartFactorial(n uint64) uint64 {
	if n <= 1 {
		return factorials[n]
	}
	if n <= uint64(len(factorials)) {
		return n * factorials[n-1]
	} else {
		res := n * smartFactorial(n-1)
		factorials = append(factorials, res)
		return res
	}
}
