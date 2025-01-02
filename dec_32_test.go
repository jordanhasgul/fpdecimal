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
			name: "encode 0 to 32-bit decimal",

			f:    float32(0),
			want: 0b0_01000_100101_00000000000000000000,
		},
		{
			name: "encode NaN to 32-bit decimal",

			f:    float32(math.NaN()),
			want: 0b0_11111_000000_11111111111111111111,
		},
		{
			name: "encode +Inf to 32-bit decimal",

			f:    float32(math.Inf(1)),
			want: 0b0_11110_110000_00000000000000000000,
		},
		{
			name: "encode -Inf to 32-bit decimal",

			f:    float32(math.Inf(-1)),
			want: 0b1_11110_110000_00000000000000000000,
		},
		{
			name: "encode -math.MaxFloat32 to 32-bit decimal",

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

func TestDec32_IsNaN(t *testing.T) {
	testCases := []struct {
		name string

		d    fpdecimal.Dec32
		want bool
	}{
		{
			name: "0 is NaN",

			d:    fpdecimal.NewDec32(0),
			want: false,
		},
		{
			name: "NaN is NaN",

			d:    fpdecimal.NewDec32(float32(math.NaN())),
			want: true,
		},
		{
			name: "+Inf is NaN",

			d:    fpdecimal.NewDec32(float32(math.Inf(1))),
			want: false,
		},
		{
			name: "-Inf is NaN",

			d:    fpdecimal.NewDec32(float32(math.Inf(-1))),
			want: false,
		},
		{
			name: "-math.MaxFloat32 is NaN",

			d:    fpdecimal.NewDec32(float32(-math.MaxFloat32)),
			want: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.d.IsNaN()
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDec32_IsInf(t *testing.T) {
	testCases := []struct {
		name string

		d    fpdecimal.Dec32
		want bool
	}{
		{
			name: "0 is +/-Inf",

			d:    fpdecimal.NewDec32(0),
			want: false,
		},
		{
			name: "NaN is +/-Inf",

			d:    fpdecimal.NewDec32(float32(math.NaN())),
			want: false,
		},
		{
			name: "+Inf is +/-Inf",

			d:    fpdecimal.NewDec32(float32(math.Inf(1))),
			want: true,
		},
		{
			name: "-Inf is +/-Inf",

			d:    fpdecimal.NewDec32(float32(math.Inf(-1))),
			want: true,
		},
		{
			name: "-math.MaxFloat32 is +/-Inf",

			d:    fpdecimal.NewDec32(float32(-math.MaxFloat32)),
			want: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.d.IsInf()
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDec32_IsPosInf(t *testing.T) {
	testCases := []struct {
		name string

		d    fpdecimal.Dec32
		want bool
	}{
		{
			name: "0 is +Inf",

			d:    fpdecimal.NewDec32(0),
			want: false,
		},
		{
			name: "NaN is +Inf",

			d:    fpdecimal.NewDec32(float32(math.NaN())),
			want: false,
		},
		{
			name: "+Inf is +Inf",

			d:    fpdecimal.NewDec32(float32(math.Inf(1))),
			want: true,
		},
		{
			name: "-Inf is +Inf",

			d:    fpdecimal.NewDec32(float32(math.Inf(-1))),
			want: false,
		},
		{
			name: "-math.MaxFloat32 is +Inf",

			d:    fpdecimal.NewDec32(float32(-math.MaxFloat32)),
			want: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.d.IsPosInf()
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDec32_IsNegInf(t *testing.T) {
	testCases := []struct {
		name string

		d    fpdecimal.Dec32
		want bool
	}{
		{
			name: "0 is -Inf",

			d:    fpdecimal.NewDec32(0),
			want: false,
		},
		{
			name: "NaN is -Inf",

			d:    fpdecimal.NewDec32(float32(math.NaN())),
			want: false,
		},
		{
			name: "+Inf is -Inf",

			d:    fpdecimal.NewDec32(float32(math.Inf(1))),
			want: false,
		},
		{
			name: "-Inf is -Inf",

			d:    fpdecimal.NewDec32(float32(math.Inf(-1))),
			want: true,
		},
		{
			name: "-math.MaxFloat32 is -Inf",

			d:    fpdecimal.NewDec32(float32(-math.MaxFloat32)),
			want: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.d.IsNegInf()
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDec32_IsZero(t *testing.T) {
	testCases := []struct {
		name string

		d    fpdecimal.Dec32
		want bool
	}{
		{
			name: "0 is 0",

			d:    fpdecimal.NewDec32(0),
			want: true,
		},
		{
			name: "NaN is 0",

			d:    fpdecimal.NewDec32(float32(math.NaN())),
			want: false,
		},
		{
			name: "+Inf is 0",

			d:    fpdecimal.NewDec32(float32(math.Inf(1))),
			want: false,
		},
		{
			name: "-Inf is 0",

			d:    fpdecimal.NewDec32(float32(math.Inf(-1))),
			want: false,
		},
		{
			name: "-math.MaxFloat32 is 0",

			d:    fpdecimal.NewDec32(float32(-math.MaxFloat32)),
			want: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.d.IsZero()
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDec32_IsPos(t *testing.T) {
	testCases := []struct {
		name string

		d    fpdecimal.Dec32
		want bool
	}{
		{
			name: "0 is positive",

			d:    fpdecimal.NewDec32(0),
			want: true,
		},
		{
			name: "NaN is positive",

			d:    fpdecimal.NewDec32(float32(math.NaN())),
			want: false,
		},
		{
			name: "+Inf is positive",

			d:    fpdecimal.NewDec32(float32(math.Inf(1))),
			want: true,
		},
		{
			name: "-Inf is positive",

			d:    fpdecimal.NewDec32(float32(math.Inf(-1))),
			want: false,
		},
		{
			name: "-math.MaxFloat32 is positive",

			d:    fpdecimal.NewDec32(float32(-math.MaxFloat32)),
			want: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.d.IsPos()
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestDec32_IsNeg(t *testing.T) {
	testCases := []struct {
		name string

		d    fpdecimal.Dec32
		want bool
	}{
		{
			name: "0 is negative",

			d:    fpdecimal.NewDec32(0),
			want: false,
		},
		{
			name: "NaN is negative",

			d:    fpdecimal.NewDec32(float32(math.NaN())),
			want: true,
		},
		{
			name: "+Inf is negative",

			d:    fpdecimal.NewDec32(float32(math.Inf(1))),
			want: false,
		},
		{
			name: "-Inf is negative",

			d:    fpdecimal.NewDec32(float32(math.Inf(-1))),
			want: true,
		},
		{
			name: "-math.MaxFloat32 is negative",

			d:    fpdecimal.NewDec32(float32(-math.MaxFloat32)),
			want: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.d.IsNeg()
			require.Equal(t, testCase.want, got)
		})
	}
}

func BenchmarkNewDec32(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for range b.N {
		_ = fpdecimal.NewDec32(-math.MaxFloat32)
	}
}
