package des_test

import (
	"testing"

	"github.com/jordanhasgul/fpdecimal/internal/des"
	"github.com/stretchr/testify/require"
)

func TestEncode32(t *testing.T) {
	testCases := []struct {
		name string

		sign        uint32
		coefficient uint32
		exponent    int32

		want uint32
	}{
		{
			name: "encode [0, 0, 0] to 32-bit decimal",

			sign:        0,
			coefficient: 0,
			exponent:    0,

			want: 0b0_01000_100101_00000000000000000000,
		},
		{
			name: "encode [1, 3402823, 32] to 32-bit decimal",

			sign:        1,
			coefficient: 3402823,
			exponent:    32,

			want: 0b1_10011_000101_10000000100100101101,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := des.Encode32(testCase.sign, testCase.coefficient, testCase.exponent)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDecode32(t *testing.T) {
	testCases := []struct {
		name string

		des uint32

		wantSign        uint32
		wantCoefficient uint32
		wantExponent    int32
	}{
		{
			name: "decode [0, 0, 0] from 32-bit decimal",

			des: 0b0_01000_100101_00000000000000000000,

			wantSign:        0,
			wantCoefficient: 0,
			wantExponent:    0,
		},
		{
			name: "decode [1, 3402823, 32] from 32-bit decimal",

			des: 0b1_10011_000101_10000000100100101101,

			wantSign:        1,
			wantCoefficient: 3402823,
			wantExponent:    32,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			gotSign, gotCoefficient, gotExponent := des.Decode32(testCase.des)
			require.Equal(t, testCase.wantSign, gotSign)
			require.Equal(t, testCase.wantCoefficient, gotCoefficient)
			require.Equal(t, testCase.wantExponent, gotExponent)
		})
	}
}

func BenchmarkEncode32(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for range b.N {
		_ = des.Encode32(1, 3402823, 32)
	}
}

func BenchmarkDecode32(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for range b.N {
		_, _, _ = des.Decode32(0b1_10011_000101_10000000100100101101)
	}
}
