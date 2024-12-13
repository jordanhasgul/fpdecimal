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
			want: 0b0_11111_000000_00000000000000000000,
		},
		{
			name: "encode +Inf as 32-bit decimal",

			f:    float32(math.Inf(1)),
			want: 0b0_11110_000000_00000000000000000000,
		},
		{
			name: "encode -Inf as 32-bit decimal",

			f:    float32(math.Inf(-1)),
			want: 0b1_11110_000000_00000000000000000000,
		},
		{
			name: "encode -1 * math.MaxFloat32 as 32-bit decimal",

			f:    -1 * float32(math.MaxFloat32),
			want: 0b1_10011_000101_10000000100100101101,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := fpdecimal.NewDec32(testCase.f)
			require.Equal(t, testCase.want, got.Bits())
		})
	}
}

func BenchmarkNewDec32(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	f := float32(math.MaxFloat32)
	for range b.N {
		_ = fpdecimal.NewDec32(f)
	}
}
