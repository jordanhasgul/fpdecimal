package bcd

const (
	MaxEncodableUint32 uint32 = 99999999
	MaxDecodableUint32 uint32 = 0b1001_1001_1001_1001_1001_1001_1001_1001

	sizeOfBCDAsUint32 uint32 = 4
	sizeOfUint32      uint32 = 32
	numBCDsInUint32          = sizeOfUint32 / sizeOfBCDAsUint32
)

func Encode32(n uint32) uint32 {
	if n > MaxEncodableUint32 {
		panic("n is greater than bcd.MaxEncodableUint32")
	}

	var bcd uint32
	for i := uint32(0); i < numBCDsInUint32; i++ {
		d := n % 10
		n /= 10

		bcd |= d << (4 * i)
	}

	return bcd
}

func Decode32(bcd uint32) uint32 {
	if bcd > MaxDecodableUint32 {
		panic("bcd is greater than bcd.MaxDecodableUint32")
	}

	var n uint32
	for i := uint32(0); i < numBCDsInUint32; i++ {
		d := bcd & 0b1111
		bcd >>= 4

		n += d * pow10[uint32](i)
	}

	return n
}

const (
	MaxEncodableUint64 uint64 = 99999999_99999999
	MaxDecodableUint64 uint64 = 0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001

	sizeOfBCDAsUint64 uint64 = 4
	sizeOfUint64      uint64 = 32
	numBCDsInUint64          = sizeOfUint64 / sizeOfBCDAsUint64
)

func Encode64(n uint64) uint64 {
	if n > MaxEncodableUint64 {
		panic("n is greater than bcd.MaxEncodableUint64")
	}

	var bcd uint64
	for i := uint64(0); i < numBCDsInUint64; i++ {
		d := n % 10
		n /= 10

		bcd |= d << (4 * i)
	}

	return bcd
}

func Decode64(bcd uint64) uint64 {
	if bcd > MaxDecodableUint64 {
		panic("n is greater than bcd.MaxDecodableUint64")
	}

	var n uint64
	for i := uint64(0); i < numBCDsInUint64; i++ {
		d := bcd & 0b1111
		bcd >>= 4

		n += d * pow10[uint64](i)
	}

	return n
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
