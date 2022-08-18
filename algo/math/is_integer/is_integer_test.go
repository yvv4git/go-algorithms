package is_integer

import (
	"fmt"
	"math"
	"testing"
)

func TestIsInt(t *testing.T) {
	type args struct {
		bits uint32
		bias int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				bits: 234523.000000,
				bias: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsInt(tt.args.bits, tt.args.bias)
			t.Log(result)
		})
	}
}

// В общем, в float32 числах:
// - старший бит - это знак(положительное или отрицательное).
// - 8 бит - это exponent.
// - 23 бита - это mantissa.
//
// В float64 вот так:
// - 1 бит - это знак.
// - 11 бит - это exponent.
// - 52 бита - это mantissa.
func TestFloatToBits(t *testing.T) {
	fmt.Printf("float32(0.1): %032b\n", math.Float32bits(0.1))   // 00111101110011001100110011001101
	fmt.Printf("float32(0.2): %032b\n", math.Float32bits(0.2))   // 00111110010011001100110011001101
	fmt.Printf("float32(-0.2): %032b\n", math.Float32bits(-0.2)) // 10111110010011001100110011001101
	fmt.Printf("float32(1): %032b\n", math.Float32bits(1))       // 00111111100000000000000000000000
	fmt.Printf("float32(1): %032b\n", math.Float32bits(-1))      // 10111111100000000000000000000000
	fmt.Printf("float32(0.1+0.1): %032b = %v \n", math.Float32bits(0.1+0.1), 0.1+0.1)
	//out1 := math.Float64bits(0)
	//out2 := math.Float64bits(1)
	//out3 := math.Float64bits(2)
	//out4 := math.Float64bits(2.5)
	//out5 := math.Float64bits(-1)
	//
	//t.Logf("Out-1: %b", out1)
	//t.Logf("Out-2: %b", out2)
	//t.Logf("Out-3: %b", out3)
	//t.Logf("Out-4: %b", out4)
	//t.Logf("Out-5: %b", out5)
}
