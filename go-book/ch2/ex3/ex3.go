package ex3

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&i)
	}
}

func PopCount(x uint64) int {
	var n int
	for i := byte(0); i < 8; i++ {
		n += int(pc[byte(x>>i*8)])
	}
	return n
}
