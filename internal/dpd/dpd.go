package dpd

import "fmt"

const (
	maxEncodableUint32 uint32 = 0b1001_1001_1001_1001_1001_1001_1001_1001
	maxDecodableUint32 uint32 = 0b0001011111_1111111101_1111111101
)

func Encode32(bcd uint32) uint32 {
	if bcd > maxEncodableUint32 {
		panicString := fmt.Sprintf("bcd is greater than %b", maxEncodableUint32)
		panic(panicString)
	}

	var (
		last3Digits      = bcd & 0b1111_1111_1111
		last3DigitsInDPD = toDPD(last3Digits)
	)
	bcd >>= 12

	var (
		next3Digits      = bcd & 0b1111_1111_1111
		next3DigitsInDPD = toDPD(next3Digits)
	)
	bcd >>= 12

	var (
		first3Digits      = bcd & 0b1111_1111_1111
		first3DigitsInDPD = toDPD(first3Digits)
	)
	bcd >>= 12

	return (first3DigitsInDPD << 20) | (next3DigitsInDPD << 10) | (last3DigitsInDPD << 0)
}

func toDPD(bcd uint32) uint32 {
	if bcd > 0b1001_1001_1001 {
		panic("bcd is greater than 0b1001_1001_1001")
	}

	var (
		a = (bcd >> 11) & 1
		b = (bcd >> 10) & 1
		c = (bcd >> 9) & 1
		d = (bcd >> 8) & 1
		e = (bcd >> 7) & 1
		f = (bcd >> 6) & 1
		g = (bcd >> 5) & 1
		h = (bcd >> 4) & 1
		i = (bcd >> 3) & 1
		j = (bcd >> 2) & 1
		k = (bcd >> 1) & 1
		l = (bcd >> 0) & 1

		q = b | (a & j) | (a & f & i)
		r = c | (a & k) | (a & g & i)
		s = d
		t = (f & (^a | ^i)) | (^a & e & j) | (e & i)
		u = g | (^a & e & k) | (a & i)
		v = h
		w = a | e | i
		x = a | (e & i) | (^e & j)
		y = e | (a & i) | (^a & k)
		z = l
	)
	return (q << 9) | (r << 8) | (s << 7) | (t << 6) | (u << 5) |
		(v << 4) | (w << 3) | (x << 2) | (y << 1) | (z << 0)
}

func Decode32(dpd uint32) uint32 {
	if dpd > maxDecodableUint32 {
		panicString := fmt.Sprintf("dpd is greater than %030b", maxDecodableUint32)
		panic(panicString)
	}

	var (
		last3DigitsInDPD = dpd & 0b1111111111
		last3Digits      = fromDPD(last3DigitsInDPD)
	)
	dpd >>= 10

	var (
		next3DigitsInDPD = dpd & 0b1111111111
		next3Digits      = fromDPD(next3DigitsInDPD)
	)
	dpd >>= 10

	var (
		first3DigitsInDPD = dpd & 0b1111111111
		first3Digits      = fromDPD(first3DigitsInDPD)
	)
	dpd >>= 10

	return (first3Digits << 24) | (next3Digits << 12) | (last3Digits << 0)
}

func fromDPD(dpd uint32) uint32 {
	if dpd > 0b1111111101 {
		panic("dpd is greater than 0b1111111101")
	}

	var (
		q = (dpd >> 9) & 1
		r = (dpd >> 8) & 1
		s = (dpd >> 7) & 1
		t = (dpd >> 6) & 1
		u = (dpd >> 5) & 1
		v = (dpd >> 4) & 1
		w = (dpd >> 3) & 1
		x = (dpd >> 2) & 1
		y = (dpd >> 1) & 1
		z = (dpd >> 0) & 1

		a = (w & x) & (^t | u | ^y)
		b = q & (^w | ^x | (t & ^u & y))
		c = r & (^w | ^x | (t & ^u & y))
		d = s
		e = w & ((^x & y) | (^u & y) | (t & y))
		f = (t & (^w | ^y)) | (q & ^t & u & w & x & y)
		g = (u & (^w | ^y)) | (r & ^t & u & x)
		h = v
		i = w & ((^x & ^y) | (x & y & (t | u)))
		j = (^w & x) | (t & w & ^x & y) | (q & x & (^y | (^t & ^u)))
		k = (^w & y) | (u & ^x & y) | (r & w & x & (^y | (^t & ^u)))
		l = z
	)
	return (a << 11) | (b << 10) | (c << 9) | (d << 8) | (e << 7) | (f << 6) |
		(g << 5) | (h << 4) | (i << 3) | (j << 2) | (k << 1) | (l << 0)
}
