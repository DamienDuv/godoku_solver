package solver


var sizeToString = map[int]string{
	2: "pair",
	3: "triple",
	4: "quadruple",
	5: "quintuple",
}

func isInArray(n int, arr []int) bool {
	for _, v := range arr {
		if n == v {
			return true
		}
	}
	return false
}
