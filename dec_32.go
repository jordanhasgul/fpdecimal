package fpdecimal

import "math"

type Dec32 struct {
	bits uint32
}

func NewDec32(f float32) *Dec32 {
	sign := uint32(0)
	if f < 0 {
		f *= -1
		sign = 1
	}

	if math.IsNaN(float64(f)) {
		return &Dec32{
			bits: (sign << signBitPosition) |
				(0b11111 << combinationFieldBitPosition),
		}
	}

	if math.IsInf(float64(f), 0) {
		return &Dec32{
			bits: (sign << signBitPosition) |
				(0b11110 << combinationFieldBitPosition),
		}
	}

	panic("unimplemented")
}

const signBitPosition = uint32(31)

func (d *Dec32) sign() uint32 {
	return d.bits >> signBitPosition
}

const (
	combinationFieldBitPosition = uint32(26)
	combinationFieldMask        = uint32(0b11111)
)

func (d *Dec32) combinationField() uint32 {
	return (d.bits >> combinationFieldBitPosition) & combinationFieldMask
}

const (
	exponentContinuationBitPosition = uint32(20)
	exponentContinuationMask        = uint32(0b111111)
)

func (d *Dec32) exponentContinuation() uint32 {
	return (d.bits >> exponentContinuationBitPosition) & exponentContinuationMask
}

const (
	coefficientContinuationBitPosition = uint32(0)
	coefficientContinuationMask        = uint32(0b11111111111111111111)
)

func (d *Dec32) coefficientContinuation() uint32 {
	return (d.bits >> coefficientContinuationBitPosition) & coefficientContinuationMask
}

func (d *Dec32) Bits() uint32 {
	return d.bits
}
