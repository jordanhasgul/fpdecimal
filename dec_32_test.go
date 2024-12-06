package fpdecimal_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/jordanhasgul/fpdecimal"
	"github.com/stretchr/testify/require"
)

func TestFromFloat32(t *testing.T) {
	testCases := []struct {
		input float32
		want  uint32
	}{
		{
			input: float32(math.NaN()),
			want:  0b0_11111_000000_00000000000000000000,
		},
		{
			input: float32(math.Inf(1)),
			want:  0b0_11110_000000_00000000000000000000,
		},
		{
			input: -7.5,
			want:  0b1_01000_100100_0000_0000_0000_0111_0101,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("encode %f to 32-bit decimal", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := fpdecimal.FromFloat32(testCase.input)
			require.Equal(t, got, testCase.want)
		})
	}
}

func TestDec32_GoString(t *testing.T) {

}

func TestDec32_String(t *testing.T) {

}
