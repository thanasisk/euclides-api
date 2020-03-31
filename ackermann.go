package main

import "errors"

func naiveAckermann(n, m uint64) (uint64, error) {
	// who likes DoS? using numbers greater than these will result to a server
	// crash owing to a stack overflow
	N_DoS_GUARD := uint64(3)
	M_DoS_GUARD := uint64(9999999)
	if n > N_DoS_GUARD || m > M_DoS_GUARD {
		// FIXME add proper message
		return 0xDEADBEEF, errors.New("too large input values for this humble CPU")
	}
	for n != 0 {
		if m == 0 {
			m = 1
		} else {
			m, _ = naiveAckermann(n, m-1) // recursive
		}
		n = n - 1
	}
	return m + 1, nil
}
