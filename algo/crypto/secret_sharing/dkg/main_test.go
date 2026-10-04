package main

import (
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const MillerRabinRounds = 64

var (
	errDrawFailed    = errors.New("draw failed")
	participantFirst = int64(FirstShare)
)

func TestNewDealer(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()

	assert.Equal(t, 1024, dealer.prime.BitLen())
	assert.Equal(t, 1023, dealer.order.BitLen())
	assert.Equal(t, big.NewInt(2), dealer.generator)
	assert.True(t, dealer.prime.ProbablyPrime(MillerRabinRounds))
	assert.True(t, dealer.order.ProbablyPrime(MillerRabinRounds))

	// Порядок 2 в группе порядка p делит q, поэтому 2^q = 1 mod p.
	power := new(big.Int).Exp(dealer.generator, dealer.order, dealer.prime)

	assert.Equal(t, big.NewInt(1), power)
}

func TestContribute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		participant int64
		threshold   int
		parts       int
		wantErr     error
	}{
		{
			name:        "first participant of three",
			participant: FirstShare,
			threshold:   DemoThreshold,
			parts:       DemoParts,
		},
		{
			name:        "last participant of three",
			participant: int64(DemoParts),
			threshold:   DemoThreshold,
			parts:       DemoParts,
		},
		{
			name:        "threshold below minimum",
			participant: FirstShare,
			threshold:   1,
			parts:       DemoParts,
			wantErr:     ErrThresholdOutOfRange,
		},
		{
			name:        "threshold above parts",
			participant: FirstShare,
			threshold:   DemoParts + 1,
			parts:       DemoParts,
			wantErr:     ErrPartsOutOfRange,
		},
		{
			name:        "parts below threshold",
			participant: FirstShare,
			threshold:   DemoThreshold,
			parts:       DemoThreshold - 1,
			wantErr:     ErrPartsOutOfRange,
		},
		{
			name:        "participant out of session",
			participant: int64(DemoParts + 1),
			threshold:   DemoThreshold,
			parts:       DemoParts,
			wantErr:     ErrPartsOutOfRange,
		},
		{
			name:        "zero participant",
			participant: 0,
			threshold:   DemoThreshold,
			parts:       DemoParts,
			wantErr:     ErrPartsOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			dealer := NewDealer()
			dealer.draw = staticDraw([]int64{1, 2, 3, 4})

			got, err := dealer.Contribute(test.participant, test.threshold, test.parts)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				assert.Empty(t, got.commitments)
				assert.Empty(t, got.shares)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, int(test.participant), got.participant)
			assert.Len(t, got.commitments, test.threshold)
			assert.Len(t, got.shares, test.parts)

			for _, commitment := range got.commitments {
				assert.True(t, commitment.Cmp(dealer.generator) >= 0)
			}
		})
	}
}

func TestContributeDrawFailure(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()
	dealer.draw = func(*big.Int) (*big.Int, error) {
		return nil, errDrawFailed
	}

	got, err := dealer.Contribute(participantFirst, DemoThreshold, DemoParts)

	require.ErrorIs(t, err, errDrawFailed)
	assert.Empty(t, got.commitments)
	assert.Empty(t, got.shares)
}

func TestVerifyShare(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()
	dealer.draw = staticDraw([]int64{4, 8})

	contribution, err := dealer.Contribute(participantFirst, DemoThreshold, DemoParts)
	require.NoError(t, err)

	forged := contribution
	forged.shares = cloneShares(contribution.shares)
	forged.shares[2] = new(big.Int).Add(contribution.shares[2], big.NewInt(1))

	tests := []struct {
		name         string
		recipient    int64
		contribution Contribution
		wantErr      error
	}{
		{name: "first recipient", recipient: FirstShare, contribution: contribution},
		{name: "last recipient", recipient: int64(DemoParts), contribution: contribution},
		{name: "own share", recipient: participantFirst, contribution: contribution},
		{
			name:         "forged share",
			recipient:    2,
			contribution: forged,
			wantErr:      ErrShareMismatch,
		},
		{
			name:         "recipient out of session",
			recipient:    int64(DemoParts + 1),
			contribution: contribution,
			wantErr:      ErrPartsOutOfRange,
		},
		{
			name:      "commitment missing",
			recipient: participantFirst,
			contribution: Contribution{
				participant: contribution.participant,
				commitments: contribution.commitments[:MinThreshold-1],
				shares:      contribution.shares,
			},
			wantErr: ErrCommitmentCount,
		},
		{
			name:      "shares missing",
			recipient: participantFirst,
			contribution: Contribution{
				participant: contribution.participant,
				commitments: contribution.commitments,
			},
			wantErr: ErrContributionMissing,
		},
		{
			name:      "participant out of session",
			recipient: participantFirst,
			contribution: Contribution{
				participant: MaxParts + 1,
				commitments: contribution.commitments,
				shares:      contribution.shares,
			},
			wantErr: ErrPartsOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := dealer.VerifyShare(test.recipient, test.contribution)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestAssemble(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()
	dealer.draw = staticDraw([]int64{3, 5, 7, 9})

	contributions, err := dealer.collect(DemoThreshold, DemoParts)
	require.NoError(t, err)

	key, err := dealer.Assemble(DemoThreshold, DemoParts, contributions)
	require.NoError(t, err)

	assert.Len(t, key.shares, DemoParts)
	assert.Equal(t, DemoThreshold, key.threshold)

	// Публичный ключ равен произведению свободных членов в экспоненциальном виде.
	expected := dealer.publicKey(contributions)
	assert.Equal(t, 0, expected.Cmp(key.value))
	assert.True(t, key.value.Cmp(big.NewInt(1)) > 0)
}

func TestAssembleRejects(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		threshold     int
		parts         int
		contributions func(dealer *Dealer) []Contribution
		wantErr       error
	}{
		{
			name:      "threshold below minimum",
			threshold: 1,
			parts:     DemoParts,
			wantErr:   ErrThresholdOutOfRange,
		},
		{
			name:      "parts below threshold",
			threshold: DemoThreshold,
			parts:     1,
			wantErr:   ErrPartsOutOfRange,
		},
		{
			name:          "no contributions",
			threshold:     DemoThreshold,
			parts:         DemoParts,
			contributions: func(*Dealer) []Contribution { return nil },
			wantErr:       ErrContributionMissing,
		},
		{
			name:      "shares of an incomplete participant",
			threshold: DemoThreshold,
			parts:     DemoParts,
			contributions: func(dealer *Dealer) []Contribution {
				contribution, err := dealer.Contribute(participantFirst, DemoThreshold, DemoParts)
				if err != nil {
					return nil
				}

				truncated := contribution
				truncated.shares = map[int]*big.Int{1: contribution.shares[1]}

				return []Contribution{truncated}
			},
			wantErr: ErrPartsOutOfRange,
		},
		{
			name:      "contribution without commitments",
			threshold: DemoThreshold,
			parts:     DemoParts,
			contributions: func(dealer *Dealer) []Contribution {
				contribution, err := dealer.Contribute(participantFirst, DemoThreshold, DemoParts)
				if err != nil {
					return nil
				}

				broken := contribution
				broken.commitments = nil

				return []Contribution{broken}
			},
			wantErr: ErrCommitmentCount,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			dealer := NewDealer()
			dealer.draw = staticDraw([]int64{3, 5, 7, 9})
			var contributions []Contribution
			if test.contributions != nil {
				contributions = test.contributions(dealer)
			}

			got, err := dealer.Assemble(test.threshold, test.parts, contributions)

			require.ErrorIs(t, err, test.wantErr)
			assert.Nil(t, got.value)
		})
	}
}

func TestCombineRestoresKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		threshold   int
		parts       int
		wantCombine bool
	}{
		{name: "two of three", threshold: 2, parts: DemoParts, wantCombine: true},
		{name: "three of three", threshold: DemoParts, parts: DemoParts, wantCombine: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			dealer := NewDealer()
			contributions, err := dealer.collect(test.threshold, test.parts)
			require.NoError(t, err)

			key, err := dealer.Assemble(test.threshold, test.parts, contributions)
			require.NoError(t, err)

			shares, err := key.Shares(test.threshold)
			require.NoError(t, err)
			require.Len(t, shares, test.threshold)

			secret, err := dealer.Combine(shares, key.value)
			require.NoError(t, err)

			// Приватный ключ согласуется с публичным: g^s = Y.
			exponentiated := new(big.Int).Exp(dealer.generator, secret, dealer.prime)
			assert.Equal(t, 0, exponentiated.Cmp(key.value))
			assert.True(t, secret.Sign() > 0)
		})
	}
}

func TestCombineAnyCombination(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()
	contributions, err := dealer.collect(DemoThreshold, DemoParts)
	require.NoError(t, err)

	key, err := dealer.Assemble(DemoThreshold, DemoParts, contributions)
	require.NoError(t, err)

	all := make([]Share, 0, DemoParts)
	for participant := FirstShare; participant <= int64(DemoParts); participant++ {
		all = append(all, Share{coordinate: participant, value: key.shares[int(participant)]})
	}

	// Любые две доли из трёх дают один и тот же приватный ключ.
	for index := range all {
		for other := index + 1; other < len(all); other++ {
			pair := []Share{all[index], all[other]}

			secret, err := dealer.Combine(pair, key.value)
			require.NoError(t, err)

			exponentiated := new(big.Int).Exp(dealer.generator, secret, dealer.prime)
			assert.Equal(t, 0, exponentiated.Cmp(key.value))
		}
	}
}

func TestCombineRejects(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()
	contributions, err := dealer.collect(DemoThreshold, DemoParts)
	require.NoError(t, err)

	key, err := dealer.Assemble(DemoThreshold, DemoParts, contributions)
	require.NoError(t, err)

	shares, err := key.Shares(DemoThreshold)
	require.NoError(t, err)

	foreign := new(big.Int).Exp(dealer.generator, big.NewInt(42), dealer.prime)

	tests := []struct {
		name    string
		shares  []Share
		public  *big.Int
		wantErr error
	}{
		{name: "no shares", shares: nil, public: key.value, wantErr: ErrNoCombinations},
		{
			name:    "repeated coordinate",
			shares:  []Share{shares[0], shares[0]},
			public:  key.value,
			wantErr: ErrRepeatedCoordinate,
		},
		{
			name:    "zero coordinate",
			shares:  []Share{{coordinate: 0, value: shares[0].value}, shares[1]},
			public:  key.value,
			wantErr: ErrCoordinateOutOfRange,
		},
		{
			name:    "value above the order",
			shares:  []Share{{coordinate: FirstShare, value: new(big.Int).Set(dealer.order)}, shares[1]},
			public:  key.value,
			wantErr: ErrValueOutOfRange,
		},
		{
			name:    "negative value",
			shares:  []Share{{coordinate: FirstShare, value: big.NewInt(-1)}, shares[1]},
			public:  key.value,
			wantErr: ErrValueOutOfRange,
		},
		{
			name:    "missing value",
			shares:  []Share{{coordinate: FirstShare}, shares[1]},
			public:  key.value,
			wantErr: ErrValueOutOfRange,
		},
		{
			name:    "foreign public key",
			shares:  shares,
			public:  foreign,
			wantErr: ErrPublicKeyMismatch,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := dealer.Combine(test.shares, test.public)

			require.ErrorIs(t, err, test.wantErr)
			assert.Nil(t, got)
		})
	}
}

func TestShares(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()
	contributions, err := dealer.collect(DemoThreshold, DemoParts)
	require.NoError(t, err)

	key, err := dealer.Assemble(DemoThreshold, DemoParts, contributions)
	require.NoError(t, err)

	tests := []struct {
		name      string
		threshold int
		wantErr   error
	}{
		{name: "threshold of the key", threshold: DemoThreshold},
		{name: "above the threshold", threshold: DemoThreshold + 1, wantErr: ErrNotEnoughShares},
		{name: "zero threshold", threshold: 0, wantErr: ErrNotEnoughShares},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := key.Shares(test.threshold)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				assert.Empty(t, got)
				return
			}

			require.NoError(t, err)
			assert.Len(t, got, test.threshold)

			for index, share := range got {
				assert.Equal(t, int64(index)+FirstShare, share.coordinate)
			}
		})
	}
}

func TestSharesMissingParticipant(t *testing.T) {
	t.Parallel()

	key := Key{shares: map[int]*big.Int{1: big.NewInt(7)}, threshold: 2}

	got, err := key.Shares(2)

	require.ErrorIs(t, err, ErrContributionMissing)
	assert.Empty(t, got)
}

func TestCommitMatchesPolynomial(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()
	coefficients := []*big.Int{big.NewInt(5), big.NewInt(3), big.NewInt(2)}
	commitments := dealer.commitments(coefficients)

	for coordinate := int64(0); coordinate <= int64(DemoParts); coordinate++ {
		got := dealer.commit(coordinate, commitments)
		want := new(big.Int).Exp(dealer.generator, dealer.evaluate(coefficients, coordinate), dealer.prime)

		assert.Equal(t, 0, want.Cmp(got), "coordinate %d", coordinate)
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()

	dealer := NewDealer()
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

			got := dealer.evaluate(coefficients, test.coordinate)

			assert.Equal(t, test.want, got.Int64())
		})
	}
}

func TestShortHex(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "2000", ShortHex(big.NewInt(8192)))
	assert.Equal(t, "0", ShortHex(big.NewInt(0)))

	huge, ok := new(big.Int).SetString(SafePrimeHex, HexBase)
	require.True(t, ok)
	assert.Len(t, ShortHex(huge), HexCharsShown)
}

func TestShareString(t *testing.T) {
	t.Parallel()

	share := Share{coordinate: FirstShare, value: big.NewInt(8192)}

	assert.Equal(t, "(1, 2000...)", share.String())
}

func TestRun(t *testing.T) {
	t.Parallel()

	assert.NoError(t, run())
}

func cloneShares(shares map[int]*big.Int) map[int]*big.Int {
	clone := make(map[int]*big.Int, len(shares))
	for participant, value := range shares {
		clone[participant] = new(big.Int).Set(value)
	}

	return clone
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
