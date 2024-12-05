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

		n += d * pow10(i)
	}

	return n
}

func pow10(n uint32) uint32 {
	if n == 0 {
		return 1
	}

	result := uint32(1)
	for i := uint32(0); i < n; i++ {
		result *= 10
	}

	return result
}
