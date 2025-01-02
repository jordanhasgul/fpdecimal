package fpdecimal

import (
	"math"

	"github.com/jordanhasgul/fpdecimal/internal/des"
)

type Dec32 struct {
	bits uint32
}

func NewDec32(f float32) Dec32 {
	f64 := float64(f)
	if math.IsNaN(f64) {
		return Dec32{
			bits: des.NaN32,
		}
	}

	if math.IsInf(f64, 1) {
		return Dec32{
			bits: des.PosInf32,
		}
	}

	if math.IsInf(f64, -1) {
		return Dec32{
			bits: des.NegInf32,
		}
	}

	if f64 == 0 {
		return Dec32{
			bits: des.Zero32,
		}
	}

	sign := uint32(0)
	if f64 < 0 {
		f64 *= -1
		sign = 1
	}

	var (
		coefficient = uint32(f64)
		exponent    = int32(0)
	)
	switch {
	case f64 >= 10000000:
		shift := math.Ceil(math.Log10(f64) - 7)
		f64 /= math.Pow10(int(shift))

		coefficient = uint32(f64)
		exponent = int32(shift)
	case f64 < 1000000:
		shift := math.Ceil(6 - math.Log10(f64))
		f64 *= math.Pow10(int(shift))

		coefficient = uint32(f64)
		exponent = int32(shift)
	}

	return Dec32{
		bits: des.Encode32(sign, coefficient, exponent),
	}
}

func (d *Dec32) IsNaN() bool {
	if d.bits == 0 {
		d.bits = des.Zero32
	}

	return d.bits == des.NaN32
}

func (d *Dec32) IsInf() bool {
	if d.bits == 0 {
		d.bits = des.Zero32
	}

	return d.IsPosInf() || d.IsNegInf()
}

func (d *Dec32) IsPosInf() bool {
	if d.bits == 0 {
		d.bits = des.Zero32
	}

	return d.bits == des.PosInf32
}

func (d *Dec32) IsNegInf() bool {
	if d.bits == 0 {
		d.bits = des.Zero32
	}

	return d.bits == des.NegInf32
}

func (d *Dec32) IsZero() bool {
	if d.bits == 0 {
		d.bits = des.Zero32
	}

	return d.bits == des.Zero32
}

func (d *Dec32) IsPos() bool {
	if d.bits == 0 {
		d.bits = des.Zero32
	}

	sign, _, _ := des.Decode32(d.bits)
	return sign == 0
}

func (d *Dec32) IsNeg() bool {
	if d.bits == 0 {
		d.bits = des.Zero32
	}

	sign, _, _ := des.Decode32(d.bits)
	return sign == 1
}

func (d *Dec32) Bits() uint32 {
	return d.bits
}
