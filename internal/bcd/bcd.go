package bcd

func Encode32(n uint32) uint32 {
	panic("unimplemented")
}

func Decode32(bcd uint32) uint32 {
	panic("unimplemented")
}

func Encode64(n uint64) uint64 {
	panic("unimplemented")
}

func Decode64(bcd uint64) uint64 {
	panic("unimplemented")
}

func pow10[N uint32 | uint64](n N) N {
	if n == 0 {
		return 1
	}

	result := N(1)
	for i := N(0); i < n; i++ {
		result *= 10
	}

	return result
}
