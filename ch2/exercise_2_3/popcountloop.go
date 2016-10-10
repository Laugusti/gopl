package popcountloop

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountLoop(x uint64) int {
	var result int
	var i uint
	for i = 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}
