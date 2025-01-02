package bcd_test

import (
	"testing"

	"github.com/jordanhasgul/fpdecimal/internal/bcd"
	"github.com/stretchr/testify/require"
)

func TestEncode32(t *testing.T) {
	t.Run("no panic on n <= 99999999", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = bcd.Encode32(99999999)
		})
	})

	t.Run("panic on n > 99999999", func(t *testing.T) {
		require.Panics(t, func() {
			_ = bcd.Encode32(99999999 + 1)
		})
	})

	testCases := []struct {
		name string

		n    uint32
		want uint32
	}{
		{
			name: "encode 0 to bcd",

			n:    0,
			want: 0b0000,
		},
		{
			name: "encode 5 to bcd",

			n:    5,
			want: 0b0101,
		},
		{
			name: "encode 11 to bcd",

			n:    11,
			want: 0b0001_0001,
		},
		{
			name: "encode 999 to bcd",

			n:    999,
			want: 0b1001_1001_1001,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := bcd.Encode32(testCase.n)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDecode32(t *testing.T) {
	t.Run("no panic on n <= 0b1001_1001_1001_1001_1001_1001_1001_1001", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = bcd.Decode32(0b1001_1001_1001_1001_1001_1001_1001_1001)
		})
	})

	t.Run("panic on n > 0b1001_1001_1001_1001_1001_1001_1001_1001", func(t *testing.T) {
		require.Panics(t, func() {
			_ = bcd.Decode32(0b1001_1001_1001_1001_1001_1001_1001_1001 + 1)
		})
	})

	testCases := []struct {
		name string

		bcd  uint32
		want uint32
	}{
		{
			name: "decode 0 from bcd",

			bcd:  0b0000,
			want: 0,
		},
		{
			name: "decode 5 from bcd",

			bcd:  0b0101,
			want: 5,
		},
		{
			name: "decode 11 from bcd",

			bcd:  0b0001_0001,
			want: 11,
		},
		{
			name: "decode 999 from bcd",

			bcd:  0b1001_1001_1001,
			want: 999,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := bcd.Decode32(testCase.bcd)
			require.Equal(t, testCase.want, got)
		})
	}
}

func BenchmarkEncode32(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for range b.N {
		_ = bcd.Encode32(99999999)
	}
}

func BenchmarkDecode32(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	
	for range b.N {
		_ = bcd.Decode32(0b1001_1001_1001_1001_1001_1001_1001_1001)
	}
}
