package ex4

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&i)
	}
}

func PopCount(x uint64) int {
	var n int
	for i := uint(0); i < 64; i++ {
		if (x>>i)&1 != 0 {
			n++
		}
	}
	return n
}
