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
		name        string
		coefficient int64
		threshold   int
		parts       int
		want        []Share
	}{
		{
			name:        "linear polynomial with two shares needed",
			coefficient: 4,
			threshold:   2,
			parts:       3,
			want:        []Share{{1, big.NewInt(17)}, {2, big.NewInt(21)}, {3, big.NewInt(25)}},
		},
		{
			name:        "quadratic polynomial with three shares needed",
			coefficient: 4,
			threshold:   3,
			parts:       3,
			want:        []Share{{1, big.NewInt(21)}, {2, big.NewInt(37)}, {3, big.NewInt(61)}},
		},
		{
			name:        "more parts than the threshold",
			coefficient: 7,
			threshold:   2,
			parts:       5,
			want: []Share{
				{1, big.NewInt(20)},
				{2, big.NewInt(27)},
				{3, big.NewInt(34)},
				{4, big.NewInt(41)},
				{5, big.NewInt(48)},
			},
		},
		{
			name:        "zero coefficients give a constant polynomial",
			coefficient: 0,
			threshold:   2,
			parts:       2,
			want:        []Share{{1, big.NewInt(13)}, {2, big.NewInt(13)}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			shamir := NewShamir(DefaultPrime)
			shamir.draw = func(*big.Int) (*big.Int, error) {
				return big.NewInt(test.coefficient), nil
			}

			got, err := shamir.Split(DemoSecret, test.threshold, test.parts)

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
			wantErr:   ErrInvalidThreshold,
		},
		{
			name:      "parts below the threshold",
			secret:    13,
			threshold: 3,
			parts:     2,
			wantErr:   ErrThresholdAboveParts,
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
		{
			name:      "parts above the field",
			secret:    13,
			threshold: 2,
			parts:     257,
			wantErr:   ErrCoordinateOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			shamir := NewShamir(DefaultPrime)
			shamir.draw = func(*big.Int) (*big.Int, error) {
				return big.NewInt(1), nil
			}

			got, err := shamir.Split(test.secret, test.threshold, test.parts)

			require.ErrorIs(t, err, test.wantErr)
			assert.Nil(t, got)
		})
	}
}

func TestSplitDrawFailure(t *testing.T) {
	t.Parallel()

	shamir := NewShamir(DefaultPrime)
	shamir.draw = func(*big.Int) (*big.Int, error) {
		return nil, errDrawFailed
	}

	got, err := shamir.Split(DemoSecret, DemoThreshold, DemoParts)

	require.ErrorIs(t, err, errDrawFailed)
	assert.Nil(t, got)
}

func TestSplitCombineRoundTrip(t *testing.T) {
	t.Parallel()

	shamir := NewShamir(DefaultPrime)

	shares, err := shamir.Split(DemoSecret, DemoThreshold, DemoParts)
	require.NoError(t, err)

	recovered, err := shamir.Combine(shares[:DemoThreshold])
	require.NoError(t, err)

	assert.Equal(t, DemoSecret, recovered.Int64())
}

func TestCombine(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []Share
		want *big.Int
	}{
		{
			name: "two shares",
			give: []Share{{1, big.NewInt(17)}, {2, big.NewInt(21)}},
			want: big.NewInt(13),
		},
		{
			name: "three shares",
			give: []Share{{1, big.NewInt(17)}, {2, big.NewInt(21)}, {3, big.NewInt(25)}},
			want: big.NewInt(13),
		},
		{
			name: "non consecutive coordinates",
			give: []Share{{1, big.NewInt(17)}, {3, big.NewInt(25)}},
			want: big.NewInt(13),
		},
		{
			name: "unordered coordinates",
			give: []Share{{3, big.NewInt(25)}, {1, big.NewInt(17)}},
			want: big.NewInt(13),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewShamir(DefaultPrime).Combine(test.give)

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
			name:    "no shares",
			give:    nil,
			wantErr: ErrNotEnoughShares,
		},
		{
			name:    "single share",
			give:    []Share{{1, big.NewInt(17)}},
			wantErr: ErrNotEnoughShares,
		},
		{
			name:    "repeated coordinate",
			give:    []Share{{1, big.NewInt(17)}, {1, big.NewInt(19)}},
			wantErr: ErrRepeatedCoordinate,
		},
		{
			name:    "coordinate zero",
			give:    []Share{{0, big.NewInt(17)}, {2, big.NewInt(21)}},
			wantErr: ErrCoordinateOutOfRange,
		},
		{
			name:    "coordinate equal to the prime",
			give:    []Share{{1, big.NewInt(17)}, {257, big.NewInt(29)}},
			wantErr: ErrCoordinateOutOfRange,
		},
		{
			name:    "value above the field",
			give:    []Share{{1, big.NewInt(17)}, {2, big.NewInt(258)}},
			wantErr: ErrValueOutOfRange,
		},
		{
			name:    "missing value",
			give:    []Share{{1, nil}, {2, big.NewInt(21)}},
			wantErr: ErrValueOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewShamir(DefaultPrime).Combine(test.give)

			require.ErrorIs(t, err, test.wantErr)
			assert.Nil(t, got)
		})
	}
}

func TestShareString(t *testing.T) {
	t.Parallel()

	share := Share{coordinate: 2, value: big.NewInt(21)}

	require.Equal(t, "(2, 21)", share.String())
}
