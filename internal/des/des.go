package des

import (
	"github.com/jordanhasgul/fpdecimal/internal/bcd"
	"github.com/jordanhasgul/fpdecimal/internal/dpd"
)

const (
	NaN32 = uint32(0b0_11111_000000_11111111111111111111)

	PosInf32 = uint32(0b0_11110_110000_00000000000000000000)
	NegInf32 = uint32(0b1_11110_110000_00000000000000000000)

	Zero32 = uint32(0b0_01000_100101_00000000000000000000)
)

func Encode32(sign, coefficient uint32, exponent int32) uint32 {
	var (
		adjustedExponent = uint32(exponent + 101)

		twoMSBsOfAdjustedExponent = (adjustedExponent & 0b11000000) >> 6
		firstDigitOfCoefficient   = (coefficient / 1e6) % 10
	)
	combinationField := 0b00000 |
		(twoMSBsOfAdjustedExponent << 3) |
		(firstDigitOfCoefficient & 0b0111)
	if firstDigitOfCoefficient >= 8 {
		combinationField = 0b11000 |
			(twoMSBsOfAdjustedExponent << 1) |
			(firstDigitOfCoefficient & 0b0001)
	}

	var (
		exponentContinuation = adjustedExponent & 0b00111111

		remainingDigitsOfCoefficient = bcd.Encode32(coefficient - (firstDigitOfCoefficient * 1e6))
		coefficientContinuation      = dpd.Encode32(remainingDigitsOfCoefficient) & 0b11111111111111111111
	)
	return (sign << 31) |
		(combinationField << 26) |
		(exponentContinuation << 20) |
		(coefficientContinuation << 0)
}

func Decode32(des uint32) (uint32, uint32, int32) {
	var (
		combinationField = (des >> 26) & 0b11111

		twoMSBsOfAdjustedExponent = (combinationField & 0b11000) >> 3
		firstDigitOfCoefficient   = combinationField & 0b00111
	)
	if (combinationField & 0b11000) == 0b11000 {
		twoMSBsOfAdjustedExponent = (combinationField & 0b00110) >> 1
		firstDigitOfCoefficient = 0b01000 | (combinationField & 0b00001)
	}

	var (
		sign = (des >> 31) & 0b1

		exponentContinuation = (des >> 20) & 0b00111111
		adjustedExponent     = (twoMSBsOfAdjustedExponent << 6) | exponentContinuation
		exponent             = int32(adjustedExponent) - 101

		coefficientContinuation      = (des >> 0) & 0b11111111111111111111
		remainingDigitsOfCoefficient = dpd.Decode32(coefficientContinuation) & 0b1111_1111_1111_1111_1111_1111
		coefficient                  = firstDigitOfCoefficient*1e6 + bcd.Decode32(remainingDigitsOfCoefficient)
	)
	return sign, coefficient, exponent
}
