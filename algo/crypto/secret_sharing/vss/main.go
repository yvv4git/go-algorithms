package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
)

// Verifiable Secret Sharing - расширение Shamir, в котором каждая доля
// проверяется участником индивидуально, до сбора порогового количества долей.
//
// Feldman VSS публикует коммитменты к коэффициентам многочлена в группе
// простого порядка:
//
//	f(x) = S + a_1·x + ... + a_{M-1}·x^(M-1)   по модулю q
//	C_k = g^(a_k)                              по модулю p
//
//	Участник проверяет свою долю: g^f(i) == Π_k C_k^(i^k) (mod p).
//	Нечестный дилер выдаёт неверную долю, и проверка это обнаруживает сразу.
const (
	// SafePrimeHex - 1024-битное простое из RFC 3526 (группа 2).
	// q = (p-1)/2 тоже простое, поэтому в группе порядка q ровно один неделимый
	// элемент, а 2 порождает подгруппу этого порядка.
	SafePrimeHex = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD1" +
		"29024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3" +
		"CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6" +
		"F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F2411" +
		"7C4B1FE649286651ECE65381FFFFFFFFFFFFFFFF"
	GeneratorHex = "02"

	HexBase = 16
)

const (
	MinShares       int   = 2  // меньше двух коэффициентов коммитменты не описывают
	FirstCoordinate int64 = 1  // доли выдаются в точках 1..N, точка 0 занята секретом
	HexCharsShown         = 16 // столько hex-символов коммитмента попадает в вывод
)

// Параметры демонстрации: секрет 13 делится на 3 доли с порогом 2,
// у каждой доли один коммитмент C_0 и верифицируются все три.
const (
	DemoSecret    int64 = 13
	DemoThreshold int   = 2
	DemoParts     int   = 3
)

// Форматы вывода вынесены в константы, чтобы vet мог проверять аргументы Printf.
const (
	FormatSecret      = "secret: %d\n"
	FormatShares      = "shares:\n"
	FormatShareLine   = "  share: %s\n"
	FormatShareShort  = "(%d, %s...)"
	FormatCommitments = "commitments:\n"
	FormatCommitment  = "  C_%d = %s\n"
	FormatVerified    = "share %s verified against commitments\n"
	FormatTampered    = "tampered share %s rejected: %v\n"
)

var (
	ErrThresholdBelowMin    = errors.New("threshold needs at least two shares")
	ErrPartsBelowThreshold  = errors.New("parts below threshold")
	ErrSecretOutOfRange     = errors.New("secret outside the group order")
	ErrCoordinateOutOfRange = errors.New("coordinate outside the field")
	ErrValueOutOfRange      = errors.New("value outside the group order")
	ErrCommitmentCount      = errors.New("unexpected number of commitments")
	ErrShareMismatch        = errors.New("share does not match the commitments")
)

// Share - доля VSS: пара (координата, значение многочлена в этой точке).
type Share struct {
	coordinate int64
	value      *big.Int
}

// String - доля в сокращённом виде: значение многочлена живёт в группе
// порядка q и занимает 1023 бита, поэтому в вывод попадает только префикс.
func (share Share) String() string {
	return fmt.Sprintf(FormatShareShort, share.coordinate, ShortHex(share.value))
}

// Distribution - результат разделения: доли участников и опубликованные
// коммитменты к коэффициентам многочлена.
type Distribution struct {
	shares      []Share
	commitments []*big.Int
}

// Feldman - проверяемое разделение секрета по схеме Feldman.
type Feldman struct {
	prime     *big.Int
	order     *big.Int
	generator *big.Int
	// draw - источник случайных коэффициентов, поле для подмены в тестах.
	draw func(limit *big.Int) (*big.Int, error)
}

// NewFeldman - схема в группе простого порядка q по модулю safe prime,
// генератор 2. Модуль и порядок берутся из констант, поэтому параметры
// схемы нельзя задать неверно.
func NewFeldman() *Feldman {
	prime, _ := new(big.Int).SetString(SafePrimeHex, HexBase)
	generator, _ := new(big.Int).SetString(GeneratorHex, HexBase)
	order := new(big.Int).Rsh(new(big.Int).Sub(prime, big.NewInt(1)), 1)
	draw := func(limit *big.Int) (*big.Int, error) { return rand.Int(rand.Reader, limit) }

	return &Feldman{prime: prime, order: order, generator: generator, draw: draw}
}

// Split - разделение секрета на parts проверяемых долей с порогом threshold.
//
//	METHOD: коэффициенты многочлена берутся по модулю q, каждый публикуется как
//	коммитмент C_k = g^(a_k) (mod p), а участник получает долю (i, f(i)).
//
//	Публикуются только коммитменты: по одному C_0 дилер раскрыть не может,
//	не решив задачу дискретного логарифма. Порядок q имеет 1023 бита, поэтому
//	координаты долей заведомо меньше него и проверяются только на неотрицательность.
//
//	TIME COMPLEXITY: O(parts * threshold) на доли плюс O(threshold) экспонент
//	на коммитменты.
func (f *Feldman) Split(secret int64, threshold, parts int) (Distribution, error) {
	if err := f.validateSplit(secret, threshold, parts); err != nil {
		return Distribution{}, err
	}

	coefficients, err := f.coefficients(secret, threshold)
	if err != nil {
		return Distribution{}, fmt.Errorf("build polynomial: %w", err)
	}

	shares := make([]Share, 0, parts)
	for index := range parts {
		coordinate := int64(index) + FirstCoordinate
		shares = append(shares, Share{coordinate: coordinate, value: f.evaluate(coefficients, coordinate)})
	}

	return Distribution{shares: shares, commitments: f.commitments(coefficients)}, nil
}

// Verify - проверка доли по опубликованным коммитментам Feldman:
//
//	g^f(i) == Π_k C_k^(i^k) (mod p).
//
//	Левая часть считается из значения доли, правая - из коммитментов дилера,
//	поэтому несовпадение означает, что доля не принадлежит многочлену.
//
//	TIME COMPLEXITY: O(threshold) модульных экспонент.
func (f *Feldman) Verify(share Share, commitments []*big.Int) error {
	if err := f.validateShare(share, commitments); err != nil {
		return err
	}

	left := new(big.Int).Exp(f.generator, share.value, f.prime)
	right := f.commit(share.coordinate, commitments)
	if left.Cmp(right) != 0 {
		return fmt.Errorf("share (%d, %s): %w", share.coordinate, ShortHex(share.value), ErrShareMismatch)
	}

	return nil
}

// coefficients - коэффициенты многочлена по модулю order, первым идёт секрет.
func (f *Feldman) coefficients(secret int64, threshold int) ([]*big.Int, error) {
	coefficients := []*big.Int{big.NewInt(secret)}

	for range threshold - 1 {
		coefficient, err := f.draw(f.order)
		if err != nil {
			return nil, fmt.Errorf("draw coefficient: %w", err)
		}

		coefficients = append(coefficients, coefficient)
	}

	return coefficients, nil
}

// commitments - коммитменты C_k = g^(a_k) по модулю prime.
func (f *Feldman) commitments(coefficients []*big.Int) []*big.Int {
	commitments := make([]*big.Int, 0, len(coefficients))
	for _, coefficient := range coefficients {
		commitment := new(big.Int).Exp(f.generator, coefficient, f.prime)
		commitments = append(commitments, commitment)
	}

	return commitments
}

// commit - произведение C_k^(x^k) по модулю prime, то есть g^f(x) для верных
// коммитментов. Показатели степени берутся по модулю order, потому что
// у элементов группы период равен order.
func (f *Feldman) commit(coordinate int64, commitments []*big.Int) *big.Int {
	result := big.NewInt(1)
	power := big.NewInt(1)
	point := big.NewInt(coordinate)

	for _, commitment := range commitments {
		term := new(big.Int).Exp(commitment, power, f.prime)
		result.Mul(result, term).Mod(result, f.prime)
		power.Mul(power, point).Mod(power, f.order)
	}

	return result
}

// evaluate - значение многочлена в точке по схеме Горнера по модулю order.
func (f *Feldman) evaluate(coefficients []*big.Int, coordinate int64) *big.Int {
	value := new(big.Int)
	point := big.NewInt(coordinate)

	for index := len(coefficients) - 1; index >= 0; index-- {
		value.Mul(value, point)
		value.Add(value, coefficients[index])
		value.Mod(value, f.order)
	}

	return value
}

// ShortHex - начало числа в hex для вывода: значения здесь занимают 1023 бита
// и не помещаются в читаемую строку, а для демонстрации хватает префикса.
func ShortHex(value *big.Int) string {
	encoded := hex.EncodeToString(value.Bytes())
	if encoded == "" {
		return "0"
	}

	return encoded[:min(len(encoded), HexCharsShown)]
}

func (f *Feldman) validateSplit(secret int64, threshold, parts int) error {
	if threshold < MinShares {
		return fmt.Errorf("threshold %d: %w", threshold, ErrThresholdBelowMin)
	}

	if parts < threshold {
		return fmt.Errorf("parts %d below threshold %d: %w", parts, threshold, ErrPartsBelowThreshold)
	}

	if secret < 0 || big.NewInt(secret).Cmp(f.order) >= 0 {
		return fmt.Errorf("secret %d: %w", secret, ErrSecretOutOfRange)
	}

	return nil
}

func (f *Feldman) validateShare(share Share, commitments []*big.Int) error {
	if share.coordinate < FirstCoordinate || big.NewInt(share.coordinate).Cmp(f.order) >= 0 {
		return fmt.Errorf("coordinate %d: %w", share.coordinate, ErrCoordinateOutOfRange)
	}

	if share.value == nil || share.value.Sign() < 0 || share.value.Cmp(f.order) >= 0 {
		return fmt.Errorf("value %v: %w", share.value, ErrValueOutOfRange)
	}

	if len(commitments) < MinShares {
		return fmt.Errorf("%d commitments: %w", len(commitments), ErrCommitmentCount)
	}

	return nil
}

// Демонстрация Feldman VSS: доли трёх участников проверяются по опубликованным
// коммитментам, а подделанная доля отвергается сразу, без сбора других долей.
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	feldman := NewFeldman()

	distribution, err := feldman.Split(DemoSecret, DemoThreshold, DemoParts)
	if err != nil {
		return fmt.Errorf("split secret: %w", err)
	}

	fmt.Printf(FormatSecret, DemoSecret)
	fmt.Print(FormatShares)
	for _, share := range distribution.shares {
		fmt.Printf(FormatShareLine, share)
	}

	return feldman.reportVerification(distribution)
}

// reportVerification - проверка всех долей и демонстрация отказа на подделке.
func (f *Feldman) reportVerification(distribution Distribution) error {
	fmt.Print(FormatCommitments)
	for index, commitment := range distribution.commitments {
		fmt.Printf(FormatCommitment, index, ShortHex(commitment))
	}

	for _, share := range distribution.shares {
		if err := f.Verify(share, distribution.commitments); err != nil {
			return fmt.Errorf("verify share: %w", err)
		}

		fmt.Printf(FormatVerified, share)
	}

	tampered := distribution.shares[0]
	tampered.value = new(big.Int).Add(tampered.value, big.NewInt(1))

	err := f.Verify(tampered, distribution.commitments)
	if !errors.Is(err, ErrShareMismatch) {
		return fmt.Errorf("verify tampered share: %w", err)
	}

	fmt.Printf(FormatTampered, tampered, err)

	return nil
}
