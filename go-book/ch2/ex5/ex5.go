package ex5

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&i)
	}
}

func PopCount(x uint64) int {
	var n int
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
