package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
)

// Additive Secret Sharing - простейшее разбиение секрета: доли складываются
// (или объединяются XOR) и дают исходный секрет.
//
//	share_1 = random(p)
//	...
//	share_{N-1} = random(p)
//	share_N = S - Σ_{i=1..N-1} share_i (mod p)
//
// Все вычисления идут в поле по модулю простого числа: без него разность
// последней доли могла бы выйти за границы поля и сумма перестала бы давать S.
const (
	DefaultPrime int64 = 257 // размер поля, секрет должен быть меньше простого
	MinParts     int   = 2   // одна доля - это сам секрет, разбиения не происходит
)

// Параметры демонстрации: секрет 13 делится на 3 доли, линейность показывается
// на секретах 7 и 5, сумма которых равна 12.
const (
	DemoSecret      int64 = 13
	DemoLeftSecret  int64 = 7
	DemoRightSecret int64 = 5
	DemoParts             = 3
)

// Форматы вывода вынесены в константы, чтобы vet мог проверять аргументы Printf.
const (
	FormatSecret    = "secret: %d\n"
	FormatShares    = "shares: %v\n"
	FormatShare     = "%d"
	FormatRecovered = "sum of %d shares: %d\n"
	FormatPartial   = "sum of %d shares is not the secret: %d\n"
	FormatLinearity = "share(%d) + share(%d) = share(%d) = %d\n"
)

var (
	ErrPartsBelowMin    = errors.New("need at least two parts")
	ErrSecretOutOfRange = errors.New("secret outside the field")
	ErrNotEnoughShares  = errors.New("no shares to combine")
	ErrValueOutOfRange  = errors.New("value outside the field")
)

// Share - доля аддитивного разделения, обычное число из поля.
// Координаты нет: сложение не привязано к точкам многочлена.
type Share struct {
	value *big.Int
}

func (share Share) String() string {
	return fmt.Sprintf(FormatShare, share.value)
}

// Additive - разбиение секрета суммированием долей над полем по модулю prime.
type Additive struct {
	prime *big.Int
	// draw - источник случайных долей, поле для подмены в тестах.
	draw func(limit *big.Int) (*big.Int, error)
}

func NewAdditive(prime int64) *Additive {
	draw := func(limit *big.Int) (*big.Int, error) { return rand.Int(rand.Reader, limit) }

	return &Additive{prime: big.NewInt(prime), draw: draw}
}

// Split - разбиение секрета на parts равных по величине долей.
//
//	METHOD: первые parts-1 долей случайны и равновероятны, последняя равна
//	S - сумме остальных. Тогда Σ share_i = S (mod p).
//
//	Случайность первых N-1 долей и есть источник perfect secrecy: зная N-1 долей,
//	получатель не может ничего сказать о секрете, потому что последняя доля
//	равномерно заполняет оставшееся множество значений.
//
//	TIME COMPLEXITY: O(parts).
func (a *Additive) Split(secret int64, parts int) ([]Share, error) {
	if err := a.validateSplit(secret, parts); err != nil {
		return nil, err
	}

	shares := make([]Share, 0, parts)
	sum := new(big.Int)

	for range parts - 1 {
		value, err := a.draw(a.prime)
		if err != nil {
			return nil, fmt.Errorf("draw share: %w", err)
		}

		shares = append(shares, Share{value: value})
		sum.Add(sum, value).Mod(sum, a.prime)
	}

	last := new(big.Int).Sub(big.NewInt(secret), sum)

	return append(shares, Share{value: last.Mod(last, a.prime)}), nil
}

// Combine - суммирование долей, результат есть секрет.
//
//	METHOD: S = Σ share_i (mod p).
//
//	Долей может быть сколько угодно, метод не знает, сколько их было при
//	разбиении, поэтому неполная сумма тоже вернётся без ошибки - это свойство
//	схемы, а не дефект проверки.
//
//	TIME COMPLEXITY: O(k), где k - количество долей.
func (a *Additive) Combine(shares []Share) (*big.Int, error) {
	if err := a.validateShares(shares); err != nil {
		return nil, err
	}

	secret := new(big.Int)
	for _, share := range shares {
		secret.Add(secret, share.value).Mod(secret, a.prime)
	}

	return secret, nil
}

// Add - сложение двух долей, гомоморфизм по сложению:
//
//	share(a) + share(b) = share(a + b).
//
//	TIME COMPLEXITY: O(1).
func (a *Additive) Add(left, right Share) Share {
	value := new(big.Int).Add(left.value, right.value)

	return Share{value: value.Mod(value, a.prime)}
}

// reportLinearity - демонстрация гомоморфизма: доли двух секретов складываются
//
//	попарно, и сумма этих долей восстанавливает сумму секретов (7 + 5 = 12).
func (a *Additive) reportLinearity() error {
	left, err := a.Split(DemoLeftSecret, DemoParts)
	if err != nil {
		return fmt.Errorf("split left secret: %w", err)
	}

	right, err := a.Split(DemoRightSecret, DemoParts)
	if err != nil {
		return fmt.Errorf("split right secret: %w", err)
	}

	sums := make([]Share, 0, DemoParts)
	for index := range DemoParts {
		sums = append(sums, a.Add(left[index], right[index]))
	}

	sum, err := a.Combine(sums)
	if err != nil {
		return fmt.Errorf("combine sums: %w", err)
	}

	fmt.Printf(FormatLinearity, DemoLeftSecret, DemoRightSecret,
		DemoLeftSecret+DemoRightSecret, sum.Int64())

	return nil
}

func (a *Additive) validateSplit(secret int64, parts int) error {
	if parts < MinParts {
		return fmt.Errorf("parts %d: %w", parts, ErrPartsBelowMin)
	}

	if secret < 0 || secret >= a.prime.Int64() {
		return fmt.Errorf("secret %d: %w", secret, ErrSecretOutOfRange)
	}

	return nil
}

func (a *Additive) validateShares(shares []Share) error {
	if len(shares) == 0 {
		return ErrNotEnoughShares
	}

	for _, share := range shares {
		if share.value == nil || share.value.Sign() < 0 || share.value.Cmp(a.prime) >= 0 {
			return fmt.Errorf("value %v: %w", share.value, ErrValueOutOfRange)
		}
	}

	return nil
}

// Демонстрация: секрет 13 делится на три равномерные доли, сумма всех трёх
//
//	восстанавливает секрет, сумма двух - нет. Затем складываются доли секретов
//	7 и 5 и восстанавливается сумма 12.
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	additive := NewAdditive(DefaultPrime)

	shares, err := additive.Split(DemoSecret, DemoParts)
	if err != nil {
		return fmt.Errorf("split secret: %w", err)
	}

	secret, err := additive.Combine(shares)
	if err != nil {
		return fmt.Errorf("combine shares: %w", err)
	}

	partial, err := additive.Combine(shares[:DemoParts-1])
	if err != nil {
		return fmt.Errorf("combine part: %w", err)
	}

	fmt.Printf(FormatSecret, DemoSecret)
	fmt.Printf(FormatShares, shares)
	fmt.Printf(FormatRecovered, DemoParts, secret.Int64())
	fmt.Printf(FormatPartial, DemoParts-1, partial.Int64())

	return additive.reportLinearity()
}
