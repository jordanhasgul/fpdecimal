package dpd

const (
	MaxEncodableUint32 uint32 = 0b1001_1001_1001_1001_1001_1001_1001_1001
	MaxDecodableUint32 uint32 = 0b0001011111_1111111101_1111111101
)

func Encode32(bcd uint32) uint32 {
	if bcd > MaxEncodableUint32 {
		panic("bcd is greater that dpd.MaxEncodableUint32")
	}

	var dpd uint32
	for i := 0; i < 3; i++ {
		first3Digits := bcd & 0b1111_1111_1111
		bcd >>= 12

		first3DigitsInDPD := toDPD(first3Digits)
		dpd |= first3DigitsInDPD << (10 * i)
	}

	return dpd
}

func Decode32(dpd uint32) uint32 {
	if dpd > MaxDecodableUint32 {
		panic("dpd is greater that dpd.MaxDecodableUint32")
	}

	var bcd uint32
	for i := 0; i < 3; i++ {
		first3DigitsInDPD := dpd & 0b1111111111
		dpd >>= 10

		first3Digits := fromDPD(first3DigitsInDPD)
		bcd |= first3Digits << (12 * i)
	}

	return bcd
}

const (
	MaxEncodableUint64 uint64 = 0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001
	MaxDecodableUint64 uint64 = 0b0001011111_1111111101_1111111101_1111111101_1111111101_1111111101
)

func Encode64(bcd uint64) uint64 {
	if bcd > MaxEncodableUint64 {
		panic("bcd is greater that dpd.MaxEncodableUint64")
	}

	var dpd uint64
	for i := 0; i < 6; i++ {
		first3Digits := bcd & 0b1111_1111_1111
		bcd >>= 12

		first3DigitsInDPD := toDPD(first3Digits)
		dpd |= first3DigitsInDPD << (10 * i)
	}

	return dpd
}

func Decode64(dpd uint64) uint64 {
	if dpd > MaxDecodableUint64 {
		panic("dpd is greater that dpd.MaxDecodableUint64")
	}

	var bcd uint64
	for i := 0; i < 6; i++ {
		first3DigitsInDPD := dpd & 0b1111111111
		dpd >>= 10

		first3Digits := fromDPD(first3DigitsInDPD)
		bcd |= first3Digits << (12 * i)
	}

	return bcd
}

func toDPD[N uint32 | uint64](bcd N) N {
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

	dpd := z << 0
	dpd |= y << 1
	dpd |= x << 2
	dpd |= w << 3
	dpd |= v << 4
	dpd |= u << 5
	dpd |= t << 6
	dpd |= s << 7
	dpd |= r << 8
	dpd |= q << 9
	return dpd
}

func fromDPD[N uint32 | uint64](dpd N) N {
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

	bcd := l << 0
	bcd |= k << 1
	bcd |= j << 2
	bcd |= i << 3
	bcd |= h << 4
	bcd |= g << 5
	bcd |= f << 6
	bcd |= e << 7
	bcd |= d << 8
	bcd |= c << 9
	bcd |= b << 10
	bcd |= a << 11
	return bcd
}
