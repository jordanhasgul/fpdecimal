package dpd_test

import (
	"fmt"
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
		input uint32
		want  uint32
	}{
		{
			input: 0b0000,
			want:  0b0000000000,
		},
		{
			input: 0b0101,
			want:  0b0000000101,
		},
		{
			input: 0b0001_0001,
			want:  0b0000010001,
		},
		{
			input: 0b1001_1001_1001,
			want:  0b0011111111,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("encode bcd(%d) to dpd", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := dpd.Encode32(testCase.input)
			require.Equal(t, got, testCase.want)
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
		input uint32
		want  uint32
	}{
		{
			input: 0b0000000000,
			want:  0b0000,
		},
		{
			input: 0b0000000101,
			want:  0b0101,
		},
		{
			input: 0b0000010001,
			want:  0b0001_0001,
		},
		{
			input: 0b0011111111,
			want:  0b1001_1001_1001,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("decode bcd(%d) from dpd", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := dpd.Decode32(testCase.input)
			require.Equal(t, got, testCase.want)
		})
	}
}

func TestEncode64(t *testing.T) {
	t.Run("no panic on n <= 0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = dpd.Encode64(0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001)
		})
	})

	t.Run("panic on n > 0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001", func(t *testing.T) {
		require.Panics(t, func() {
			_ = dpd.Encode64(0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001 + 1)
		})
	})

	testCases := []struct {
		input uint64
		want  uint64
	}{
		{
			input: 0b0000,
			want:  0b0000000000,
		},
		{
			input: 0b0101,
			want:  0b0000000101,
		},
		{
			input: 0b0001_0001,
			want:  0b0000010001,
		},
		{
			input: 0b1001_1001_1001,
			want:  0b0011111111,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("encode bcd(%d) to dpd", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := dpd.Encode64(testCase.input)
			require.Equal(t, got, testCase.want)
		})
	}
}

func TestDecode64(t *testing.T) {
	t.Run("no panic on n <= 0b0001011111_1111111101_1111111101_1111111101_1111111101_1111111101", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = dpd.Decode64(0b0001011111_1111111101_1111111101_1111111101_1111111101_1111111101)
		})
	})

	t.Run("panic on n > 0b0001011111_1111111101_1111111101_1111111101_1111111101_1111111101", func(t *testing.T) {
		require.Panics(t, func() {
			_ = dpd.Decode64(0b0001011111_1111111101_1111111101_1111111101_1111111101_1111111101 + 1)
		})
	})

	testCases := []struct {
		input uint64
		want  uint64
	}{
		{
			input: 0b0000000000,
			want:  0b0000,
		},
		{
			input: 0b0000000101,
			want:  0b0101,
		},
		{
			input: 0b0000010001,
			want:  0b0001_0001,
		},
		{
			input: 0b0011111111,
			want:  0b1001_1001_1001,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("decode bcd(%d) from dpd", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := dpd.Decode64(testCase.input)
			require.Equal(t, got, testCase.want)
		})
	}
}
