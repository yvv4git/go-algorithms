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
		name  string
		drawn []int64
		want  []Share
	}{
		{
			name:  "last share closes the sum to the secret",
			drawn: []int64{5, 20},
			want:  []Share{{big.NewInt(5)}, {big.NewInt(20)}, {big.NewInt(245)}},
		},
		{
			name:  "zero drawn shares",
			drawn: []int64{0, 0},
			want:  []Share{{big.NewInt(0)}, {big.NewInt(0)}, {big.NewInt(13)}},
		},
		{
			name:  "sum wraps around the field",
			drawn: []int64{200, 100},
			want:  []Share{{big.NewInt(200)}, {big.NewInt(100)}, {big.NewInt(227)}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			additive := NewAdditive(DefaultPrime)
			additive.draw = staticDraw(test.drawn)

			got, err := additive.Split(DemoSecret, DemoParts)

			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestSplitRejects(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		secret  int64
		parts   int
		wantErr error
	}{
		{
			name:    "single part",
			secret:  13,
			parts:   1,
			wantErr: ErrPartsBelowMin,
		},
		{
			name:    "secret equal to the prime",
			secret:  DefaultPrime,
			parts:   3,
			wantErr: ErrSecretOutOfRange,
		},
		{
			name:    "negative secret",
			secret:  -1,
			parts:   3,
			wantErr: ErrSecretOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			additive := NewAdditive(DefaultPrime)
			additive.draw = staticDraw([]int64{1, 2})

			got, err := additive.Split(test.secret, test.parts)

			require.ErrorIs(t, err, test.wantErr)
			assert.Nil(t, got)
		})
	}
}

func TestSplitDrawFailure(t *testing.T) {
	t.Parallel()

	additive := NewAdditive(DefaultPrime)
	additive.draw = func(*big.Int) (*big.Int, error) {
		return nil, errDrawFailed
	}

	got, err := additive.Split(DemoSecret, DemoParts)

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
			name: "three shares",
			give: []Share{{big.NewInt(5)}, {big.NewInt(20)}, {big.NewInt(245)}},
			want: big.NewInt(13),
		},
		{
			name: "two shares of three",
			give: []Share{{big.NewInt(5)}, {big.NewInt(20)}},
			want: big.NewInt(25),
		},
		{
			name: "single share",
			give: []Share{{big.NewInt(13)}},
			want: big.NewInt(13),
		},
		{
			name: "sum wraps around the field",
			give: []Share{{big.NewInt(200)}, {big.NewInt(100)}},
			want: big.NewInt(43),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewAdditive(DefaultPrime).Combine(test.give)

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
			name:    "value equal to the prime",
			give:    []Share{{big.NewInt(13)}, {big.NewInt(257)}},
			wantErr: ErrValueOutOfRange,
		},
		{
			name:    "missing value",
			give:    []Share{{big.NewInt(13)}, {nil}},
			wantErr: ErrValueOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewAdditive(DefaultPrime).Combine(test.give)

			require.ErrorIs(t, err, test.wantErr)
			assert.Nil(t, got)
		})
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []Share
		want Share
	}{
		{
			name: "plain addition",
			give: []Share{{big.NewInt(5)}, {big.NewInt(20)}},
			want: Share{big.NewInt(25)},
		},
		{
			name: "addition wraps around the field",
			give: []Share{{big.NewInt(200)}, {big.NewInt(100)}},
			want: Share{big.NewInt(43)},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := NewAdditive(DefaultPrime).Add(test.give[0], test.give[1])

			assert.Equal(t, test.want, got)
		})
	}
}

func TestSplitCombineRoundTrip(t *testing.T) {
	t.Parallel()

	additive := NewAdditive(DefaultPrime)

	shares, err := additive.Split(DemoSecret, DemoParts)
	require.NoError(t, err)

	recovered, err := additive.Combine(shares)
	require.NoError(t, err)

	assert.Equal(t, DemoSecret, recovered.Int64())
}

func TestShareString(t *testing.T) {
	t.Parallel()

	share := Share{big.NewInt(42)}

	require.Equal(t, "42", share.String())
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
