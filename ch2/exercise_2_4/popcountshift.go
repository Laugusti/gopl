package popcountshift

func PopCountShift(x uint64) int {
	var result int
	for i := 0; i < 64; i++ {
		result += int(x & 1)
		x = x >> 1
	}
	return result
}
