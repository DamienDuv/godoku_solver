package models

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

// Bits indexes return the 0 base index of every '1' in 
// the base 2 representation of n 
func BitsIndexes(n uint16) []int {
	arr := []int{}
	i := 0

	for n > 0 {
		if n & 1 == 1{
			arr = append(arr, i)
		}
		i++
		n >>=1
	}

	return arr
}