package main

// IsolatedBitIndex checks if a uint16 is composed of only a single 1
// (i.e n is a power of 2) and returns either (-1, false) if it isn't
// otherwise it's base 1 index and true. (log2(n), true)
func IsolatedBitIndex(n uint16) (int, bool) {
	if n == 0 || n & (n-1) != 0 {
		return -1, false
	}

	res := 0
	for n > 0 {
		n >>= 1
		res++
	}

	return res, true
}