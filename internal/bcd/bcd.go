package bcd

import (
	"fmt"
)

const (
	maxEncodableUint32 uint32 = 99999999
	maxDecodableUint32 uint32 = 0b1001_1001_1001_1001_1001_1001_1001_1001
)

func Encode32(n uint32) uint32 {
	if n > maxEncodableUint32 {
		panicString := fmt.Sprintf("n is greater than %d", maxEncodableUint32)
		panic(panicString)
	}

	digit8 := n % 10
	n /= 10

	digit7 := n % 10
	n /= 10

	digit6 := n % 10
	n /= 10

	digit5 := n % 10
	n /= 10

	digit4 := n % 10
	n /= 10

	digit3 := n % 10
	n /= 10

	digit2 := n % 10
	n /= 10

	digit1 := n % 10
	n /= 10

	return (digit1 << 28) | (digit2 << 24) | (digit3 << 20) | (digit4 << 16) |
		(digit5 << 12) | (digit6 << 8) | (digit7 << 4) | (digit8 << 0)
}

func Decode32(bcd uint32) uint32 {
	if bcd > maxDecodableUint32 {
		panicString := fmt.Sprintf("bcd is greater than %b", maxDecodableUint32)
		panic(panicString)
	}

	digit8 := bcd & 0b1111
	bcd >>= 4

	digit7 := bcd & 0b1111
	bcd >>= 4

	digit6 := bcd & 0b1111
	bcd >>= 4

	digit5 := bcd & 0b1111
	bcd >>= 4

	digit4 := bcd & 0b1111
	bcd >>= 4

	digit3 := bcd & 0b1111
	bcd >>= 4

	digit2 := bcd & 0b1111
	bcd >>= 4

	digit1 := bcd & 0b1111
	bcd >>= 4

	return (digit1 * 10000000) + (digit2 * 1000000) + (digit3 * 100000) + (digit4 * 10000) +
		(digit5 * 1000) + (digit6 * 100) + (digit7 * 10) + (digit8 * 1)
}
