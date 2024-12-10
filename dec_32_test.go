package fpdecimal_test

import (
	"math"
	"testing"

	"github.com/jordanhasgul/fpdecimal"
	"github.com/stretchr/testify/require"
)

func TestNewDec32(t *testing.T) {
	testCases := []struct {
		name string

		f    float32
		want uint32
	}{
		{
			name: "encode NaN as 32-bit decimal",

			f:    float32(math.NaN()),
			want: 0b0_11110_000000_00000000000000000000,
		},
		{
			name: "encode +Inf as 32-bit decimal",

			f:    float32(math.Inf(1)),
			want: 0b0_11111_000000_00000000000000000000,
		},
		{
			name: "encode -Inf as 32-bit decimal",

			f:    float32(math.Inf(-1)),
			want: 0b1_11111_000000_00000000000000000000,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := fpdecimal.NewDec32(testCase.f)
			require.Equal(t, testCase.want, got.Bits())
		})
	}
}
