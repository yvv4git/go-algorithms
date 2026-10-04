package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
)

// Distributed Key Generation - совместная генерация общего публичного ключа,
// при которой никто из участников не знает соответствующего приватного ключа.
//
// Joint-Feldman DKG (Gennaro, Jarecki, Krawczyk, Rabin, 1999) состоит из трёх фаз:
//
//  1. Каждый участник i выбирает многочлен f_i степени t-1 со свободным
//     членом a_{i0} и публикует коммитменты C_{ik} = g^(a_{ik}) (mod p),
//     а затем отправляет свою долю f_i(j) каждому участнику j.
//  2. Участник j проверяет полученную долю через коммитменты участника i:
//     g^(f_i(j)) == Π_k C_{ik}^(j^k) (mod p), иначе публикует complaint.
//  3. Общий публичный ключ Y = Π_i C_{i0}, приватный ключ s = Σ_i a_{i0}
//     никому неизвестен, а своя доля участника j равна Σ_i f_i(j).
//
// Дилера нет: уйти может любой участник, и остальные собирают ключ из своих долей.
const (
	// SafePrimeHex - 1024-битное простое из RFC 3526 (группа 2).
	// q = (p-1)/2 тоже простое, поэтому 2 порождает подгруппу этого порядка.
	SafePrimeHex = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD1" +
		"29024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3" +
		"CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6" +
		"F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F2411" +
		"7C4B1FE649286651ECE65381FFFFFFFFFFFFFFFF"
	GeneratorHex = "02"

	HexBase = 16
)

const (
	MinThreshold  int   = 2  // порог меньше двух не собирает ключ
	MaxThreshold  int   = 9  // порог не может превышать число участников сессии
	MaxParts      int   = 9  // столько участников допускается в сессии
	FirstShare    int64 = 1  // доли нумеруются с 1, точка 0 занята приватным ключом
	HexCharsShown       = 16 // столько hex-символов значения попадает в вывод
)

// Параметры демонстрации: три участника собирают общий ключ с порогом 2,
// каждый проверяет доли остальных и публикует по два коммитмента.
const (
	DemoThreshold int = 2
	DemoParts     int = 3
)

// Форматы вывода вынесены в константы, чтобы vet мог проверять аргументы Printf.
const (
	FormatSession     = "participants: %d, threshold: %d\n"
	FormatCommitments = "contributions:\n"
	FormatCommitment  = "  participant %d publishes C_0 = %s\n"
	FormatVerified    = "participant %d verified share from participant %d\n"
	FormatPublicKey   = "public key: %s\n"
	FormatShares      = "shares of the key:\n"
	FormatShareLine   = "  share: %s\n"
	FormatShareShort  = "(%d, %s...)"
	FormatSecretKey   = "assembled secret key: %s\n"
	FormatRejected    = "forged share of participant %d is rejected: %v\n"
)

var (
	ErrThresholdOutOfRange  = errors.New("threshold outside the allowed range")
	ErrPartsOutOfRange      = errors.New("participants outside the allowed range")
	ErrContributionMissing  = errors.New("contribution missing")
	ErrCommitmentCount      = errors.New("unexpected number of commitments")
	ErrCoordinateOutOfRange = errors.New("coordinate outside the field")
	ErrValueOutOfRange      = errors.New("value outside the group order")
	ErrRepeatedCoordinate   = errors.New("repeated coordinate among shares")
	ErrNotInvertible        = errors.New("denominator is not invertible")
	ErrShareMismatch        = errors.New("share does not match the commitments")
	ErrPublicKeyMismatch    = errors.New("reconstructed key differs from the public key")
	ErrNoCombinations       = errors.New("no combination of shares collected")
	ErrNotEnoughShares      = errors.New("not enough shares for the threshold")
)

// Share - доля приватного ключа: пара (координата, значение многочлена).
type Share struct {
	coordinate int64
	value      *big.Int
}

// String - доля в сокращённом виде: значения живут в группе порядка q,
// поэтому в вывод попадает только префикс.
func (share Share) String() string {
	return fmt.Sprintf(FormatShareShort, share.coordinate, ShortHex(share.value))
}

// Contribution - вклад участника сессии: коммитменты к его коэффициентам
// и доли, которые он раздал остальным участникам.
type Contribution struct {
	participant int
	commitments []*big.Int
	shares      map[int]*big.Int
}

// Key - результат генерации: общий публичный ключ, доли участников и порог.
type Key struct {
	value     *big.Int
	shares    map[int]*big.Int
	threshold int
}

// Dealer - Joint-Feldman DKG в группе простого порядка q по модулю safe prime.
type Dealer struct {
	prime     *big.Int
	order     *big.Int
	generator *big.Int
	// draw - источник случайных коэффициентов, поле для подмены в тестах.
	draw func(limit *big.Int) (*big.Int, error)
}

// NewDealer - сессия DKG с параметрами из констант: модуль, порядок группы
// и генератор 2. Поэтому параметры группы нельзя задать неверно.
func NewDealer() *Dealer {
	prime, _ := new(big.Int).SetString(SafePrimeHex, HexBase)
	generator, _ := new(big.Int).SetString(GeneratorHex, HexBase)
	order := new(big.Int).Rsh(new(big.Int).Sub(prime, big.NewInt(1)), 1)
	draw := func(limit *big.Int) (*big.Int, error) { return rand.Int(rand.Reader, limit) }

	return &Dealer{prime: prime, order: order, generator: generator, draw: draw}
}

// Contribute - фаза 1: участник participant публикует коммитменты к коэффициентам
// своего многочлена и раздаёт доли f_i(j) всем участникам сессии.
//
//	METHOD: свободный член a_{i0} и коэффициенты a_{i1} ... a_{i,t-1} берутся
//	по модулю q, публикуются коммитменты C_{ik} = g^(a_{ik}) (mod p). Значения
//	долей остаются у отправителя: он отдаёт их участникам лично.
//
//	TIME COMPLEXITY: O(parts * threshold) на доли плюс O(threshold) экспонент.
func (d *Dealer) Contribute(participant int64, threshold, parts int) (Contribution, error) {
	if err := d.validateSession(threshold, parts); err != nil {
		return Contribution{}, err
	}

	if !d.validParticipant(participant, parts) {
		return Contribution{}, fmt.Errorf("participant %d: %w", participant, ErrPartsOutOfRange)
	}

	coefficients, err := d.coefficients(threshold)
	if err != nil {
		return Contribution{}, fmt.Errorf("draw coefficients: %w", err)
	}

	return Contribution{
		participant: int(participant),
		commitments: d.commitments(coefficients),
		shares:      d.distribute(coefficients, parts),
	}, nil
}

// VerifyShare - фаза 2: участник recipient проверяет полученную от contributor
// долю по его коммитментам Feldman:
//
//	g^(f_c(j)) == Π_k C_{ck}^(j^k) (mod p).
//
//	Несовпадение означает, что contributor прислал неверную долю, и участник
//	публикует complaint вместо того, чтобы принять плохие данные.
//
//	TIME COMPLEXITY: O(threshold) модульных экспонент.
func (d *Dealer) VerifyShare(recipient int64, contribution Contribution) error {
	if err := d.validateContribution(contribution); err != nil {
		return err
	}

	if !d.validParticipant(recipient, len(contribution.shares)) {
		return fmt.Errorf("recipient %d: %w", recipient, ErrPartsOutOfRange)
	}

	value, ok := contribution.shares[int(recipient)]
	if !ok {
		return fmt.Errorf("recipient %d: %w", recipient, ErrContributionMissing)
	}

	left := new(big.Int).Exp(d.generator, value, d.prime)
	right := d.commit(recipient, contribution.commitments)
	if left.Cmp(right) != 0 {
		return fmt.Errorf("recipient %d from %d: %w", recipient, contribution.participant, ErrShareMismatch)
	}

	return nil
}

// Assemble - фаза 3: проверка вкладов всех участников и сбор общего
// публичного ключа вместе с долей каждого участника.
//
//	METHOD: публичный ключ Y = Π_i C_{i0}, а доля участника j равна
//	Σ_i f_i(j) по модулю q. Приватный ключ s = Σ_i a_{i0} не знает никто,
//	поэтому никто не может и подписать единолично.
//
//	TIME COMPLEXITY: O(parts^2 * threshold) на проверку и сбор долей.
func (d *Dealer) Assemble(threshold, parts int, contributions []Contribution) (Key, error) {
	if err := d.validateSession(threshold, parts); err != nil {
		return Key{}, err
	}

	if err := d.verifyContributions(parts, contributions); err != nil {
		return Key{}, err
	}

	return Key{
		value:     d.publicKey(contributions),
		shares:    d.shares(contributions),
		threshold: threshold,
	}, nil
}

// Shares - доли первых threshold участников для восстановления приватного ключа.
//
//	METHOD: координаты идут по порядку номеров участников, а значения берутся
//	из уже собранного ключа, поэтому собранный набор можно передать в Combine.
func (key Key) Shares(threshold int) ([]Share, error) {
	if threshold <= 0 || threshold > key.threshold {
		return nil, fmt.Errorf("threshold %d of %d: %w", threshold, key.threshold, ErrNotEnoughShares)
	}

	shares := make([]Share, 0, threshold)
	for participant := FirstShare; len(shares) < threshold; participant++ {
		value, ok := key.shares[int(participant)]
		if !ok {
			return nil, fmt.Errorf("participant %d: %w", participant, ErrContributionMissing)
		}

		shares = append(shares, Share{coordinate: participant, value: value})
	}

	return shares, nil
}

// Combine - восстановление приватного ключа по threshold долям и проверка
// результата по публичному ключу.
//
//	METHOD: интерполяция Лагранжа в точке 0, где живёт приватный ключ:
//
//	s = f(0) = Σ y_i · L_i(0), где L_i(0) = Π (0 - x_j) / (x_i - x_j).
//
//	Каждая доля - сумма многочленов всех участников, поэтому полученный
//	многочлен имеет степень не выше threshold-1 и любой набор из threshold
//	долей даёт один и тот же ключ.
//
//	TIME COMPLEXITY: O(threshold^2).
func (d *Dealer) Combine(shares []Share, public *big.Int) (*big.Int, error) {
	if len(shares) == 0 {
		return nil, ErrNoCombinations
	}

	if err := d.validateShares(shares); err != nil {
		return nil, err
	}

	secret, err := d.interpolate(shares)
	if err != nil {
		return nil, fmt.Errorf("interpolate shares: %w", err)
	}

	if expected := new(big.Int).Exp(d.generator, secret, d.prime); expected.Cmp(public) != 0 {
		return nil, fmt.Errorf("got %s, %w", ShortHex(secret), ErrPublicKeyMismatch)
	}

	return secret, nil
}

// collect - фазы 1 и 2 для всех участников сессии: каждый вносит свои
// коэффициенты и проверяет доли остальных участников.
func (d *Dealer) collect(threshold, parts int) ([]Contribution, error) {
	contributions := make([]Contribution, 0, parts)

	for participant := FirstShare; participant <= int64(parts); participant++ {
		contribution, err := d.Contribute(participant, threshold, parts)
		if err != nil {
			return nil, fmt.Errorf("participant %d: %w", participant, err)
		}

		contributions = append(contributions, contribution)
	}

	return contributions, d.verifyAll(contributions)
}

// verifyAll - каждый участник проверяет доли, полученные от остальных.
func (d *Dealer) verifyAll(contributions []Contribution) error {
	for _, contribution := range contributions {
		for recipient := FirstShare; recipient <= int64(len(contribution.shares)); recipient++ {
			if err := d.VerifyShare(recipient, contribution); err != nil {
				return err
			}
		}
	}

	return nil
}

// coefficients - коэффициенты многочлена участника по модулю order,
// свободный член идёт первым и публикуется только коммитментом.
func (d *Dealer) coefficients(threshold int) ([]*big.Int, error) {
	coefficients := make([]*big.Int, 0, threshold)

	for range threshold {
		coefficient, err := d.draw(d.order)
		if err != nil {
			return nil, fmt.Errorf("draw coefficient: %w", err)
		}

		coefficients = append(coefficients, coefficient)
	}

	return coefficients, nil
}

// commitments - коммитменты C_k = g^(a_k) по модулю prime.
func (d *Dealer) commitments(coefficients []*big.Int) []*big.Int {
	commitments := make([]*big.Int, 0, len(coefficients))
	for _, coefficient := range coefficients {
		commitment := new(big.Int).Exp(d.generator, coefficient, d.prime)
		commitments = append(commitments, commitment)
	}

	return commitments
}

// distribute - значение многочлена участника в точке каждого получателя.
func (d *Dealer) distribute(coefficients []*big.Int, parts int) map[int]*big.Int {
	shares := make(map[int]*big.Int, parts)

	for recipient := FirstShare; recipient <= int64(parts); recipient++ {
		shares[int(recipient)] = d.evaluate(coefficients, recipient)
	}

	return shares
}

// evaluate - значение многочлена в точке по схеме Горнера по модулю order.
func (d *Dealer) evaluate(coefficients []*big.Int, coordinate int64) *big.Int {
	value := new(big.Int)
	point := new(big.Int).Mod(big.NewInt(coordinate), d.order)

	for index := len(coefficients) - 1; index >= 0; index-- {
		value.Mul(value, point)
		value.Add(value, coefficients[index])
		value.Mod(value, d.order)
	}

	return value
}

// commit - произведение C_k^(x^k) по модулю prime, то есть g^f(x) для верных
// коммитментов. Показатели степени берутся по модулю order, потому что
// у элементов группы период равен order.
func (d *Dealer) commit(coordinate int64, commitments []*big.Int) *big.Int {
	result := big.NewInt(1)
	power := big.NewInt(1)
	point := new(big.Int).Mod(big.NewInt(coordinate), d.order)

	for _, commitment := range commitments {
		term := new(big.Int).Exp(commitment, power, d.prime)
		result.Mul(result, term).Mod(result, d.prime)
		power.Mul(power, point).Mod(power, d.order)
	}

	return result
}

// interpolate - значение многочлена в точке 0 по формуле Лагранжа.
func (d *Dealer) interpolate(shares []Share) (*big.Int, error) {
	secret := big.NewInt(0)

	for index, share := range shares {
		weight, err := d.lagrange(shares, index)
		if err != nil {
			return nil, err
		}

		secret.Add(secret, new(big.Int).Mul(share.value, weight))
		secret.Mod(secret, d.order)
	}

	return secret, nil
}

// lagrange - базисный полином Лагранжа для доли с индексом index в точке 0:
// произведение (0 - x_j)/(x_i - x_j) по всем остальным долям.
func (d *Dealer) lagrange(shares []Share, index int) (*big.Int, error) {
	numerator := big.NewInt(1)
	denominator := big.NewInt(1)
	own := new(big.Int).Mod(big.NewInt(shares[index].coordinate), d.order)

	for position, share := range shares {
		if position == index {
			continue
		}

		point := new(big.Int).Mod(big.NewInt(share.coordinate), d.order)
		numerator.Mul(numerator, new(big.Int).Sub(d.order, point))
		numerator.Mod(numerator, d.order)
		denominator.Mul(denominator, new(big.Int).Sub(own, point))
		denominator.Mod(denominator, d.order)
	}

	inverse := new(big.Int).ModInverse(denominator, d.order)
	if inverse == nil {
		return nil, fmt.Errorf("denominator %s: %w", ShortHex(denominator), ErrNotInvertible)
	}

	return numerator.Mul(numerator, inverse).Mod(numerator, d.order), nil
}

// verifyContributions - вклады проверяются по форме: у каждого должны быть
// коммитменты и доли для всех участников сессии.
func (d *Dealer) verifyContributions(parts int, contributions []Contribution) error {
	if len(contributions) == 0 {
		return ErrContributionMissing
	}

	for _, contribution := range contributions {
		if err := d.validateContribution(contribution); err != nil {
			return err
		}

		if len(contribution.shares) != parts {
			return fmt.Errorf("participant %d shares %d of %d: %w",
				contribution.participant, len(contribution.shares), parts, ErrPartsOutOfRange)
		}
	}

	return nil
}

// publicKey - произведение свободных членов в экспоненциальном виде:
//
//	Y = Π_i g^(a_{i0}) = Π_i C_{i0}.
func (d *Dealer) publicKey(contributions []Contribution) *big.Int {
	key := big.NewInt(1)

	for _, contribution := range contributions {
		key.Mul(key, contribution.commitments[0])
		key.Mod(key, d.prime)
	}

	return key
}

// shares - доли участников: сумма по модулю order значений всех многочленов.
func (d *Dealer) shares(contributions []Contribution) map[int]*big.Int {
	shares := make(map[int]*big.Int)

	for _, contribution := range contributions {
		for recipient, value := range contribution.shares {
			sum, ok := shares[recipient]
			if !ok {
				sum = big.NewInt(0)
			}

			sum.Add(sum, value)
			shares[recipient] = sum.Mod(sum, d.order)
		}
	}

	return shares
}

func (d *Dealer) validParticipant(participant int64, parts int) bool {
	return participant >= FirstShare && participant <= int64(parts)
}

func (d *Dealer) validateSession(threshold, parts int) error {
	if threshold < MinThreshold || threshold > MaxThreshold {
		return fmt.Errorf("threshold %d: %w", threshold, ErrThresholdOutOfRange)
	}

	if parts < threshold || parts > MaxParts {
		return fmt.Errorf("parts %d: %w", parts, ErrPartsOutOfRange)
	}

	return nil
}

func (d *Dealer) validateContribution(contribution Contribution) error {
	if !d.validParticipant(int64(contribution.participant), MaxParts) {
		return fmt.Errorf("participant %d: %w", contribution.participant, ErrPartsOutOfRange)
	}

	if len(contribution.commitments) < MinThreshold {
		return fmt.Errorf("%d commitments: %w", len(contribution.commitments), ErrCommitmentCount)
	}

	if len(contribution.shares) == 0 {
		return fmt.Errorf("participant %d: %w", contribution.participant, ErrContributionMissing)
	}

	return nil
}

func (d *Dealer) validateShares(shares []Share) error {
	seen := make(map[int64]struct{}, len(shares))

	for _, share := range shares {
		if share.coordinate < FirstShare || new(big.Int).Mod(big.NewInt(share.coordinate), d.order).Sign() == 0 {
			return fmt.Errorf("coordinate %d: %w", share.coordinate, ErrCoordinateOutOfRange)
		}

		if share.value == nil || share.value.Sign() < 0 || share.value.Cmp(d.order) >= 0 {
			return fmt.Errorf("value %v: %w", share.value, ErrValueOutOfRange)
		}

		if _, ok := seen[share.coordinate]; ok {
			return fmt.Errorf("coordinate %d: %w", share.coordinate, ErrRepeatedCoordinate)
		}

		seen[share.coordinate] = struct{}{}
	}

	return nil
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

// Демонстрация Joint-Feldman DKG: три участника вносят свои коэффициенты,
// проверяют доли друг друга, собирают общий публичный ключ и восстанавливают
// приватный ключ по двум долям.
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	dealer := NewDealer()

	contributions, err := dealer.collect(DemoThreshold, DemoParts)
	if err != nil {
		return fmt.Errorf("collect contributions: %w", err)
	}

	key, err := dealer.Assemble(DemoThreshold, DemoParts, contributions)
	if err != nil {
		return fmt.Errorf("assemble key: %w", err)
	}

	return dealer.reportKey(key, contributions)
}

// reportKey - вывод собранного ключа, его долей и восстановленного секрета.
func (d *Dealer) reportKey(key Key, contributions []Contribution) error {
	fmt.Printf(FormatSession, DemoParts, DemoThreshold)
	fmt.Print(FormatCommitments)
	for _, contribution := range contributions {
		fmt.Printf(FormatCommitment, contribution.participant, ShortHex(contribution.commitments[0]))
	}

	fmt.Printf(FormatPublicKey, ShortHex(key.value))

	shares, err := key.Shares(DemoThreshold)
	if err != nil {
		return fmt.Errorf("take shares: %w", err)
	}

	fmt.Print(FormatShares)
	for _, share := range shares {
		fmt.Printf(FormatShareLine, share)
	}

	secret, err := d.Combine(shares, key.value)
	if err != nil {
		return fmt.Errorf("combine shares: %w", err)
	}

	fmt.Printf(FormatSecretKey, ShortHex(secret))

	return d.reportComplaint(contributions[0])
}

// reportComplaint - нечестный участник подменяет долю, и проверка сразу
// отклоняет её: плохие данные не попадают в агрегирование.
func (d *Dealer) reportComplaint(contribution Contribution) error {
	recipient := FirstShare
	forged := contribution
	forged.shares = make(map[int]*big.Int, len(contribution.shares))
	for participant, value := range contribution.shares {
		forged.shares[participant] = new(big.Int).Set(value)
	}

	recipientShare := forged.shares[int(recipient)]
	forged.shares[int(recipient)] = new(big.Int).Add(recipientShare, big.NewInt(1))

	err := d.VerifyShare(recipient, forged)
	if !errors.Is(err, ErrShareMismatch) {
		return fmt.Errorf("verify forged share: %w", err)
	}

	fmt.Printf(FormatRejected, contribution.participant, err)

	return nil
}
