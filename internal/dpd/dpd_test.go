package dpd_test

import (
	"testing"

	"github.com/jordanhasgul/fpdecimal/internal/dpd"
	"github.com/stretchr/testify/require"
)

func TestEncode32(t *testing.T) {
	t.Run("no panic on n <= 0b1001_1001_1001_1001_1001_1001_1001_1001", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = dpd.Encode32(0b1001_1001_1001_1001_1001_1001_1001_1001)
		})
	})

	t.Run("panic on n > 0b1001_1001_1001_1001_1001_1001_1001_1001", func(t *testing.T) {
		require.Panics(t, func() {
			_ = dpd.Encode32(0b1001_1001_1001_1001_1001_1001_1001_1001 + 1)
		})
	})

	testCases := []struct {
		name string

		bcd  uint32
		want uint32
	}{
		{
			name: "encode bcd(0000) to dpd",

			bcd:  0b0000,
			want: 0b0000000000,
		},
		{
			name: "encode bcd(0101) to dpd",

			bcd:  0b0101,
			want: 0b0000000101,
		},
		{
			name: "encode bcd(0001_0001) to dpd",

			bcd:  0b0001_0001,
			want: 0b0000010001,
		},
		{
			name: "encode bcd(1001_1001_1001) to dpd",

			bcd:  0b1001_1001_1001,
			want: 0b0011111111,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := dpd.Encode32(testCase.bcd)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDecode32(t *testing.T) {
	t.Run("no panic on n <= 0b0001011111_1111111101_1111111101", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = dpd.Decode32(0b0001011111_1111111101_1111111101)
		})
	})

	t.Run("panic on n > 0b0001011111_1111111101_1111111101", func(t *testing.T) {
		require.Panics(t, func() {
			_ = dpd.Decode32(0b0001011111_1111111101_1111111101 + 1)
		})
	})

	testCases := []struct {
		name string

		dpd  uint32
		want uint32
	}{
		{
			name: "decode bcd(0000) from dpd",

			dpd:  0b0000000000,
			want: 0b0000,
		},
		{
			name: "decode bcd(0101) from dpd",

			dpd:  0b0000000101,
			want: 0b0101,
		},
		{
			name: "decode bcd(0001_0001) from dpd",

			dpd:  0b0000010001,
			want: 0b0001_0001,
		},
		{
			name: "decode bcd(1001_1001_1001) from dpd",

			dpd:  0b0011111111,
			want: 0b1001_1001_1001,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := dpd.Decode32(testCase.dpd)
			require.Equal(t, testCase.want, got)
		})
	}
}
