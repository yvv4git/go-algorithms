package main

import (
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errDrawFailed = errors.New("draw failed")

func TestSplit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		drawn     []int64
		threshold int
		want      []Share
	}{
		{
			name:      "planes through one point",
			drawn:     []int64{2, 3, 5, 7},
			threshold: 2,
			want: []Share{
				{[]*big.Int{big.NewInt(3)}, big.NewInt(19)},
				{[]*big.Int{big.NewInt(5)}, big.NewInt(23)},
				{[]*big.Int{big.NewInt(7)}, big.NewInt(27)},
			},
		},
		{
			name:      "threshold three needs three unknowns",
			drawn:     []int64{2, 3, 5, 7, 11, 13, 17, 20},
			threshold: 3,
			want: []Share{
				{[]*big.Int{big.NewInt(5), big.NewInt(7)}, big.NewInt(44)},
				{[]*big.Int{big.NewInt(11), big.NewInt(13)}, big.NewInt(74)},
				{[]*big.Int{big.NewInt(17), big.NewInt(20)}, big.NewInt(107)},
			},
		},
		{
			name:      "zero point and plane coefficients",
			drawn:     []int64{0, 0, 0, 0},
			threshold: 2,
			want: []Share{
				{[]*big.Int{big.NewInt(0)}, big.NewInt(13)},
				{[]*big.Int{big.NewInt(0)}, big.NewInt(13)},
				{[]*big.Int{big.NewInt(0)}, big.NewInt(13)},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			blakley := NewBlakley(DefaultPrime)
			blakley.draw = staticDraw(test.drawn)

			got, err := blakley.Split(DemoSecret, test.threshold, DemoParts)

			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestSplitRejects(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		secret    int64
		threshold int
		parts     int
		wantErr   error
	}{
		{
			name:      "threshold below the minimum",
			secret:    13,
			threshold: 1,
			parts:     3,
			wantErr:   ErrThresholdBelowMin,
		},
		{
			name:      "parts below the threshold",
			secret:    13,
			threshold: 3,
			parts:     2,
			wantErr:   ErrPartsBelowThreshold,
		},
		{
			name:      "secret equal to the prime",
			secret:    DefaultPrime,
			threshold: 2,
			parts:     3,
			wantErr:   ErrSecretOutOfRange,
		},
		{
			name:      "negative secret",
			secret:    -1,
			threshold: 2,
			parts:     3,
			wantErr:   ErrSecretOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			blakley := NewBlakley(DefaultPrime)
			blakley.draw = staticDraw([]int64{1, 2, 3, 4, 5, 6})

			got, err := blakley.Split(test.secret, test.threshold, test.parts)

			require.ErrorIs(t, err, test.wantErr)
			assert.Nil(t, got)
		})
	}
}

func TestSplitDrawFailure(t *testing.T) {
	t.Parallel()

	blakley := NewBlakley(DefaultPrime)
	blakley.draw = func(*big.Int) (*big.Int, error) {
		return nil, errDrawFailed
	}

	got, err := blakley.Split(DemoSecret, DemoThreshold, DemoParts)

	require.ErrorIs(t, err, errDrawFailed)
	assert.Nil(t, got)
}

func TestCombine(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []Share
		want *big.Int
	}{
		{
			name: "two planes",
			give: []Share{
				{[]*big.Int{big.NewInt(3)}, big.NewInt(19)},
				{[]*big.Int{big.NewInt(5)}, big.NewInt(23)},
			},
			want: big.NewInt(13),
		},
		{
			name: "unordered planes",
			give: []Share{
				{[]*big.Int{big.NewInt(5)}, big.NewInt(23)},
				{[]*big.Int{big.NewInt(3)}, big.NewInt(19)},
			},
			want: big.NewInt(13),
		},
		{
			name: "three planes with two coefficients",
			give: []Share{
				{[]*big.Int{big.NewInt(5), big.NewInt(7)}, big.NewInt(44)},
				{[]*big.Int{big.NewInt(11), big.NewInt(13)}, big.NewInt(74)},
				{[]*big.Int{big.NewInt(17), big.NewInt(20)}, big.NewInt(107)},
			},
			want: big.NewInt(13),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewBlakley(DefaultPrime).Combine(test.give)

			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestCombineRejects(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		give    []Share
		wantErr error
	}{
		{
			name:    "no planes",
			give:    nil,
			wantErr: ErrNotEnoughShares,
		},
		{
			name: "single plane",
			give: []Share{
				{[]*big.Int{big.NewInt(3)}, big.NewInt(19)},
			},
			wantErr: ErrNotEnoughShares,
		},
		{
			name: "fewer planes than unknowns",
			give: []Share{
				{[]*big.Int{big.NewInt(5), big.NewInt(7)}, big.NewInt(44)},
				{[]*big.Int{big.NewInt(11), big.NewInt(13)}, big.NewInt(74)},
			},
			wantErr: ErrNotEnoughShares,
		},
		{
			name: "planes of different arity",
			give: []Share{
				{[]*big.Int{big.NewInt(3)}, big.NewInt(19)},
				{[]*big.Int{big.NewInt(5), big.NewInt(7)}, big.NewInt(44)},
			},
			wantErr: ErrShareArity,
		},
		{
			name: "plane without coefficients",
			give: []Share{
				{nil, big.NewInt(19)},
				{[]*big.Int{big.NewInt(5)}, big.NewInt(23)},
			},
			wantErr: ErrShareArity,
		},
		{
			name: "parallel planes",
			give: []Share{
				{[]*big.Int{big.NewInt(3)}, big.NewInt(19)},
				{[]*big.Int{big.NewInt(3)}, big.NewInt(25)},
			},
			wantErr: ErrSingularSystem,
		},
		{
			name: "value equal to the prime",
			give: []Share{
				{[]*big.Int{big.NewInt(3)}, big.NewInt(19)},
				{[]*big.Int{big.NewInt(5)}, big.NewInt(257)},
			},
			wantErr: ErrValueOutOfRange,
		},
		{
			name: "coefficient equal to the prime",
			give: []Share{
				{[]*big.Int{big.NewInt(257)}, big.NewInt(19)},
				{[]*big.Int{big.NewInt(5)}, big.NewInt(23)},
			},
			wantErr: ErrValueOutOfRange,
		},
		{
			name: "missing value",
			give: []Share{
				{[]*big.Int{big.NewInt(3)}, nil},
				{[]*big.Int{big.NewInt(5)}, big.NewInt(23)},
			},
			wantErr: ErrValueOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewBlakley(DefaultPrime).Combine(test.give)

			require.ErrorIs(t, err, test.wantErr)
			assert.Nil(t, got)
		})
	}
}

func TestSplitCombineRoundTrip(t *testing.T) {
	t.Parallel()

	blakley := NewBlakley(DefaultPrime)

	shares, err := blakley.Split(DemoSecret, DemoThreshold, DemoParts)
	require.NoError(t, err)

	recovered, err := blakley.Combine(shares[:DemoThreshold])
	require.NoError(t, err)

	assert.Equal(t, DemoSecret, recovered.Int64())
}

func TestShareString(t *testing.T) {
	t.Parallel()

	share := Share{[]*big.Int{big.NewInt(3), big.NewInt(5)}, big.NewInt(19)}

	require.Equal(t, "y = x0 + 3·x1 + 5·x2 = 19", share.String())
}

func staticDraw(values []int64) func(limit *big.Int) (*big.Int, error) {
	drawn := make([]*big.Int, 0, len(values))
	for _, value := range values {
		drawn = append(drawn, big.NewInt(value))
	}

	return func(*big.Int) (*big.Int, error) {
		next := drawn[0]
		drawn = drawn[1:]

		return next, nil
	}
}
