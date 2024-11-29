package bcd_test

import (
	"fmt"
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
		input uint32
		want  uint32
	}{
		{
			input: 0,
			want:  0b0000,
		},
		{
			input: 5,
			want:  0b0101,
		},
		{
			input: 11,
			want:  0b0001_0001,
		},
		{
			input: 999,
			want:  0b1001_1001_1001,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("encode %d to bcd", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := bcd.Encode32(testCase.input)
			require.Equal(t, got, testCase.want)
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
		input uint32
		want  uint32
	}{
		{
			input: 0b0000,
			want:  0,
		},
		{
			input: 0b0101,
			want:  5,
		},
		{
			input: 0b0001_0001,
			want:  11,
		},
		{
			input: 0b1001_1001_1001,
			want:  999,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("decode %d from bcd", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := bcd.Decode32(testCase.input)
			require.Equal(t, got, testCase.want)
		})
	}
}

func TestEncode64(t *testing.T) {
	t.Run("no panic on n <= 99999999_99999999", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = bcd.Encode64(99999999_99999999)
		})
	})

	t.Run("panic on n > 99999999_99999999", func(t *testing.T) {
		require.Panics(t, func() {
			_ = bcd.Encode64(99999999_99999999 + 1)
		})
	})

	testCases := []struct {
		input uint64
		want  uint64
	}{
		{
			input: 0,
			want:  0b0000,
		},
		{
			input: 5,
			want:  0b0101,
		},
		{
			input: 11,
			want:  0b0001_0001,
		},
		{
			input: 999,
			want:  0b1001_1001_1001,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("encode %d to bcd", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := bcd.Encode64(testCase.input)
			require.Equal(t, got, testCase.want)
		})
	}
}

func TestDecode64(t *testing.T) {
	t.Run("no panic on n <= 0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = bcd.Decode64(0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001)
		})
	})

	t.Run("panic on n > 0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001", func(t *testing.T) {
		require.Panics(t, func() {
			_ = bcd.Decode64(0b1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001_1001 + 1)
		})
	})

	testCases := []struct {
		name string

		input uint64
		want  uint64
	}{
		{
			input: 0b0000,
			want:  0,
		},
		{
			input: 0b0101,
			want:  5,
		},
		{
			input: 0b0001_0001,
			want:  11,
		},
		{
			input: 0b1001_1001_1001,
			want:  999,
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("decode %d from bcd", testCase.input)
		t.Run(name, func(t *testing.T) {
			got := bcd.Decode64(testCase.input)
			require.Equal(t, got, testCase.want)
		})
	}
}
