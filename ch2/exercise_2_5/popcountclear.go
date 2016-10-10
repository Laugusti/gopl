package popcountclear

func PopCountClear(x uint64) int {
	result := 0
	for x != 0 {
		result++
		x &= x - 1
	}
	return result
}
