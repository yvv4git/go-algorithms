package main

import (
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const MillerRabinRounds = 64

var errDrawFailed = errors.New("draw failed")

func TestNewFeldman(t *testing.T) {
	t.Parallel()

	feldman := NewFeldman()

	assert.Equal(t, 1024, feldman.prime.BitLen())
	assert.Equal(t, 1023, feldman.order.BitLen())
	assert.Equal(t, big.NewInt(2), feldman.generator)
	assert.True(t, feldman.prime.ProbablyPrime(MillerRabinRounds))
	assert.True(t, feldman.order.ProbablyPrime(MillerRabinRounds))

	// Порядок 2 в группе порядка p делит q, поэтому 2^q = 1 mod p.
	power := new(big.Int).Exp(feldman.generator, feldman.order, feldman.prime)

	assert.Equal(t, big.NewInt(1), power)
}

func TestSplit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		secret    int64
		threshold int
		parts     int
	}{
		{name: "threshold two of three", secret: DemoSecret, threshold: 2, parts: DemoParts},
		{name: "threshold equals parts", secret: DemoSecret, threshold: DemoParts, parts: DemoParts},
		{name: "secret zero", secret: 0, threshold: DemoThreshold, parts: DemoParts},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			feldman := NewFeldman()
			feldman.draw = staticDraw([]int64{1, 2, 3, 4})

			got, err := feldman.Split(test.secret, test.threshold, test.parts)

			require.NoError(t, err)
			assert.Len(t, got.shares, test.parts)
			assert.Len(t, got.commitments, test.threshold)

			for index, share := range got.shares {
				assert.Equal(t, int64(index)+FirstCoordinate, share.coordinate)
				require.NoError(t, feldman.Verify(share, got.commitments))
			}
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
			name:      "threshold below minimum",
			secret:    DemoSecret,
			threshold: 1,
			parts:     DemoParts,
			wantErr:   ErrThresholdBelowMin,
		},
		{
			name:      "parts below threshold",
			secret:    DemoSecret,
			threshold: DemoParts,
			parts:     2,
			wantErr:   ErrPartsBelowThreshold,
		},
		{
			name:      "negative secret",
			secret:    -1,
			threshold: DemoThreshold,
			parts:     DemoParts,
			wantErr:   ErrSecretOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			feldman := NewFeldman()

			got, err := feldman.Split(test.secret, test.threshold, test.parts)

			require.ErrorIs(t, err, test.wantErr)
			assert.Empty(t, got.shares)
			assert.Empty(t, got.commitments)
		})
	}
}

func TestSplitDrawFailure(t *testing.T) {
	t.Parallel()

	feldman := NewFeldman()
	feldman.draw = func(*big.Int) (*big.Int, error) {
		return nil, errDrawFailed
	}

	got, err := feldman.Split(DemoSecret, DemoThreshold, DemoParts)

	require.ErrorIs(t, err, errDrawFailed)
	assert.Empty(t, got.shares)
}

func TestSplitHidesSecret(t *testing.T) {
	t.Parallel()

	feldman := NewFeldman()
	feldman.draw = staticDraw([]int64{5})

	distribution, err := feldman.Split(DemoSecret, DemoThreshold, DemoParts)
	require.NoError(t, err)

	// В точке 0 произведение коммитментов равно C_0 = g^S: секрет зашит
	// в свободный член и в доли не попадает.
	atZero := feldman.commit(0, distribution.commitments)
	want := new(big.Int).Exp(feldman.generator, big.NewInt(DemoSecret), feldman.prime)

	assert.Equal(t, 0, want.Cmp(atZero))

	for _, share := range distribution.shares {
		assert.NotEqual(t, DemoSecret, share.value.Int64())
	}
}

func TestVerify(t *testing.T) {
	t.Parallel()

	feldman := NewFeldman()
	feldman.draw = staticDraw([]int64{9})

	distribution, err := feldman.Split(DemoSecret, DemoThreshold, DemoParts)
	require.NoError(t, err)

	tamperedValue := new(big.Int).Add(distribution.shares[0].value, big.NewInt(1))
	brokenCommitment := new(big.Int).Mul(distribution.commitments[1], feldman.generator)
	brokenCommitment.Mod(brokenCommitment, feldman.prime)

	tests := []struct {
		name        string
		share       Share
		commitments []*big.Int
		wantErr     error
	}{
		{
			name:        "first share",
			share:       distribution.shares[0],
			commitments: distribution.commitments,
		},
		{
			name:        "last share",
			share:       distribution.shares[DemoParts-1],
			commitments: distribution.commitments,
		},
		{
			name:        "tampered value",
			share:       Share{coordinate: distribution.shares[0].coordinate, value: tamperedValue},
			commitments: distribution.commitments,
			wantErr:     ErrShareMismatch,
		},
		{
			name: "tampered coordinate",
			share: Share{
				coordinate: distribution.shares[0].coordinate + 1,
				value:      distribution.shares[0].value,
			},
			commitments: distribution.commitments,
			wantErr:     ErrShareMismatch,
		},
		{
			name:        "forged commitment",
			share:       distribution.shares[0],
			commitments: []*big.Int{distribution.commitments[0], brokenCommitment},
			wantErr:     ErrShareMismatch,
		},
		{
			name:        "commitment missing",
			share:       distribution.shares[0],
			commitments: distribution.commitments[:MinShares-1],
			wantErr:     ErrCommitmentCount,
		},
		{
			name:        "zero coordinate",
			share:       Share{coordinate: 0, value: big.NewInt(DemoSecret)},
			commitments: distribution.commitments,
			wantErr:     ErrCoordinateOutOfRange,
		},
		{
			name:        "value above the order",
			share:       Share{coordinate: FirstCoordinate, value: new(big.Int).Set(feldman.order)},
			commitments: distribution.commitments,
			wantErr:     ErrValueOutOfRange,
		},
		{
			name:        "negative value",
			share:       Share{coordinate: FirstCoordinate, value: big.NewInt(-1)},
			commitments: distribution.commitments,
			wantErr:     ErrValueOutOfRange,
		},
		{
			name:        "missing value",
			share:       Share{coordinate: FirstCoordinate},
			commitments: distribution.commitments,
			wantErr:     ErrValueOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := feldman.Verify(test.share, test.commitments)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestCommitMatchesPolynomial(t *testing.T) {
	t.Parallel()

	feldman := NewFeldman()
	coefficients := []*big.Int{big.NewInt(5), big.NewInt(3), big.NewInt(2)}
	commitments := feldman.commitments(coefficients)

	for coordinate := int64(0); coordinate <= int64(DemoParts); coordinate++ {
		got := feldman.commit(coordinate, commitments)
		want := new(big.Int).Exp(feldman.generator, feldman.evaluate(coefficients, coordinate), feldman.prime)

		assert.Equal(t, 0, want.Cmp(got), "coordinate %d", coordinate)
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()

	feldman := NewFeldman()
	coefficients := []*big.Int{big.NewInt(5), big.NewInt(3), big.NewInt(2)}
	tests := []struct {
		name       string
		coordinate int64
		want       int64
	}{
		{name: "free term", coordinate: 0, want: 5},
		{name: "first coefficient", coordinate: 1, want: 10},
		{name: "square", coordinate: 2, want: 19},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := feldman.evaluate(coefficients, test.coordinate)

			assert.Equal(t, test.want, got.Int64())
		})
	}
}

func TestShortHex(t *testing.T) {
	t.Parallel()

	short := big.NewInt(8192)
	assert.Equal(t, "2000", ShortHex(short))
	assert.Equal(t, "0", ShortHex(big.NewInt(0)))

	huge, ok := new(big.Int).SetString(SafePrimeHex, HexBase)
	require.True(t, ok)
	assert.Len(t, ShortHex(huge), HexCharsShown)
}

func TestShareString(t *testing.T) {
	t.Parallel()

	share := Share{coordinate: FirstCoordinate, value: big.NewInt(8192)}

	assert.Equal(t, "(1, 2000...)", share.String())
}

func TestRun(t *testing.T) {
	t.Parallel()

	assert.NoError(t, run())
}

func staticDraw(values []int64) func(limit *big.Int) (*big.Int, error) {
	index := 0

	return func(limit *big.Int) (*big.Int, error) {
		value := big.NewInt(values[index])
		if index < len(values)-1 {
			index++
		}

		if value.Sign() < 0 || value.Cmp(limit) >= 0 {
			return nil, errDrawFailed
		}

		return value, nil
	}
}
