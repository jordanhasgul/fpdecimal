package fpdecimal

import (
	"math"

	"github.com/jordanhasgul/fpdecimal/internal/bcd"
	"github.com/jordanhasgul/fpdecimal/internal/dpd"
)

const (
	signBitPosition                    = uint32(31)
	combinationFieldBitPosition        = uint32(26)
	combinationFieldMask               = uint32(0b11111)
	exponentBias                       = int32(101)
	exponentContinuationBitPosition    = uint32(20)
	exponentContinuationMask           = uint32(0b111111)
	coefficientContinuationBitPosition = uint32(0)
	coefficientContinuationMask        = uint32(0b11111111111111111111)
)

type Dec32 struct {
	bits uint32
}

func NewDec32(f float32) Dec32 {
	sign := uint32(0)
	if f < 0 {
		f *= -1
		sign = 1
	}

	if math.IsNaN(float64(f)) {
		return Dec32{
			bits: (sign << signBitPosition) |
				(0b11111 << combinationFieldBitPosition),
		}
	}

	if math.IsInf(float64(f), 0) {
		return Dec32{
			bits: (sign << signBitPosition) |
				(0b11110 << combinationFieldBitPosition),
		}
	}

	var (
		coefficient, exponent = extractCoefficientAndExponent(f)

		adjustedExponent          = uint32(exponent + exponentBias)
		twoMSBsOfAdjustedExponent = adjustedExponent >> 6

		firstDigitOfCoefficient = (coefficient / 1e6) % 10
		remainingDigits         = coefficient - (firstDigitOfCoefficient * 1e6)

		combinationField        = constructCombinationField(firstDigitOfCoefficient, twoMSBsOfAdjustedExponent)
		exponentContinuation    = constructExponentContinuation(adjustedExponent)
		coefficientContinuation = constructCoefficientContinuation(remainingDigits)
	)
	return Dec32{
		bits: (sign << signBitPosition) |
			(combinationField << combinationFieldBitPosition) |
			(exponentContinuation << exponentContinuationBitPosition) |
			(coefficientContinuation << coefficientContinuationBitPosition),
	}
}

func extractCoefficientAndExponent(f float32) (uint32, int32) {
	exponent := int32(0)
	switch {
	case f < 1000000:
		for f < 1000000 {
			f *= 10
			exponent--
		}
	case f >= 10000000:
		for f >= 10000000 {
			f /= 10
			exponent++
		}
	}

	coefficient := uint32(f)
	return coefficient, exponent
}

func constructCombinationField(firstDigitOfCoefficient, twoMSBsOfAdjustedExponent uint32) uint32 {
	var combinationField uint32
	if firstDigitOfCoefficient >= 8 {
		combinationField = 0b11000 |
			(twoMSBsOfAdjustedExponent << 1) |
			(firstDigitOfCoefficient & 0b0001)
	} else {
		combinationField = 0b00000 |
			(twoMSBsOfAdjustedExponent << 3) |
			(firstDigitOfCoefficient & 0b0111)
	}

	return combinationField & combinationFieldMask
}

func constructCoefficientContinuation(remainingDigits uint32) uint32 {
	remainingDigits = bcd.Encode32(remainingDigits)
	return dpd.Encode32(remainingDigits) & coefficientContinuationMask
}

func constructExponentContinuation(adjustedExponent uint32) uint32 {
	return adjustedExponent & exponentContinuationMask
}

func (d *Dec32) Bits() uint32 {
	return d.bits
}
