package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
)

// Схема разделения секрета Шмира: секрет разбивается на N долей, любые M из N
// восстанавливают его, а M-1 не дают никакой информации о секрете.
//
// Все вычисления идут в конечном поле по модулю простого числа: именно это
// гарантирует, что при восстановлении деление на ненулевые элементы всегда
// определено, а координаты долей остаются различимыми.
const (
	DefaultPrime    int64 = 257 // размер поля, секрет должен быть меньше простого
	MinShares       int   = 2   // меньше двух точек многочлен не восстановить
	FirstCoordinate int64 = 1   // доли выдаются в точках 1..N, точка 0 занята секретом
	Origin          int64 = 0   // секрет это f(0), поэтому координата 0 долям не раздаётся
)

// Параметры демонстрации: секрет 13 делится на 3 доли с порогом 2,
// то есть восстановить его могут любые две доли из трёх.
const (
	DemoSecret    int64 = 13
	DemoThreshold int   = 2
	DemoParts     int   = 3
)

// Форматы вывода вынесены в константы, чтобы vet мог проверять аргументы Printf.
const (
	FormatSecret         = "secret: %d\n"
	FormatShares         = "shares: %v\n"
	FormatShare          = "(%d, %v)"
	FormatRecovered      = "recovered from %d shares: %d\n"
	FormatBelowThreshold = "%d shares below the threshold are rejected\n"
)

var (
	ErrInvalidThreshold     = errors.New("threshold needs at least two shares")
	ErrThresholdAboveParts  = errors.New("parts below threshold")
	ErrSecretOutOfRange     = errors.New("secret outside the field")
	ErrNotEnoughShares      = errors.New("not enough shares")
	ErrCoordinateOutOfRange = errors.New("coordinate outside the field")
	ErrValueOutOfRange      = errors.New("value outside the field")
	ErrRepeatedCoordinate   = errors.New("repeated coordinate")
)

// Share - одна доля: пара (координата, значение многочлена в этой координате).
type Share struct {
	coordinate int64
	value      *big.Int
}

func (share Share) String() string {
	return fmt.Sprintf(FormatShare, share.coordinate, share.value)
}

// Shamir - разделение секрета над полем по модулю простого числа prime.
type Shamir struct {
	prime *big.Int
	// draw - источник случайных коэффициентов многочлена, поле для подмены в тестах.
	draw func(limit *big.Int) (*big.Int, error)
}

func NewShamir(prime int64) *Shamir {
	draw := func(limit *big.Int) (*big.Int, error) { return rand.Int(rand.Reader, limit) }

	return &Shamir{prime: big.NewInt(prime), draw: draw}
}

// Split - разделение секрета на parts долей с порогом threshold.
//
//	METHOD: строится многочлен
//	f(x) = S + a_1·x + a_2·x^2 + ... + a_{M-1}·x^(M-1),
//	где S = secret - свободный член, a_i - случайные коэффициенты по модулю prime.
//	Доля i - это значение f(i). Секрет спрятан в точке 0, участникам не выдаётся.
//
//	Степень многочлена = threshold - 1, поэтому любые threshold точек однозначно
//	задают многочлен, а threshold-1 точек оставляют threshold-1 неизвестных
//	коэффициентов - это и даёт perfect secrecy.
//
//	TIME COMPLEXITY: O(parts * threshold).
func (s *Shamir) Split(secret int64, threshold, parts int) ([]Share, error) {
	if err := s.validateSplit(secret, threshold, parts); err != nil {
		return nil, err
	}

	// Свободный член - это секрет, остальные коэффициенты случайны.
	coefficients, err := s.coefficients(secret, threshold-1)
	if err != nil {
		return nil, fmt.Errorf("build polynomial: %w", err)
	}

	shares := make([]Share, 0, parts)
	for index := range parts {
		coordinate := int64(index) + FirstCoordinate
		value := s.evaluate(coefficients, coordinate)

		shares = append(shares, Share{coordinate: coordinate, value: value})
	}

	return shares, nil
}

// Combine - восстановление секрета из переданных долей, их должно быть не меньше порога.
//
//	METHOD: интерполяция Лагранжа, вычисленная сразу в точке 0:
//
//		S = f(0) = Σ y_i · L_i(0),   где L_i(0) = Π (0 - x_j) / (x_i - x_j)
//
//	Знаменатель обращается модульной инверсией: он ненулевой, так как координаты
//	долей различны. Результат не зависит от того, какие именно доли передали и в
//	каком порядке, поэтому лишние доли просто добавляют слагаемые.
//
//	TIME COMPLEXITY: O(k^2), где k - количество переданных долей.
func (s *Shamir) Combine(shares []Share) (*big.Int, error) {
	if err := s.validateShares(shares); err != nil {
		return nil, err
	}

	// Сумма y_i · L_i(0): каждый член верен только в своей точке и равен нулю в чужих.
	secret := new(big.Int)
	for index, share := range shares {
		basis, err := s.lagrange(shares, index)
		if err != nil {
			return nil, err
		}

		secret.Add(secret, new(big.Int).Mul(share.value, basis))
	}

	return secret.Mod(secret, s.prime), nil
}

// coefficients - коэффициенты многочлена от старшего к свободному члену,
//
//	первым идёт секрет, остальные - случайные значения из поля.
func (s *Shamir) coefficients(secret int64, degree int) ([]*big.Int, error) {
	coefficients := []*big.Int{big.NewInt(secret)}

	for range degree {
		coefficient, err := s.draw(s.prime)
		if err != nil {
			return nil, fmt.Errorf("draw coefficient: %w", err)
		}

		coefficients = append(coefficients, coefficient)
	}

	return coefficients, nil
}

// evaluate - значение многочлена в точке coordinate по схеме Горнера:
//
//	f(x) = (...((a_k · x + a_{k-1}) · x + ...) · x + a_0)
//
//	Схема Горнера даёт O(degree) умножений вместо O(degree^2), если считать
//	степени наивным возведением в степень.
func (s *Shamir) evaluate(coefficients []*big.Int, coordinate int64) *big.Int {
	value := new(big.Int)
	point := big.NewInt(coordinate)

	for index := len(coefficients) - 1; index >= 0; index-- {
		value.Mul(value, point)
		value.Add(value, coefficients[index])
		value.Mod(value, s.prime)
	}

	return value
}

// lagrange - базис Лагранжа L_i(0) для доли shares[index]: произведение числителя
//
//	на обратный знаменатель, где оба произведения берутся по всем остальным долям.
func (s *Shamir) lagrange(shares []Share, index int) (*big.Int, error) {
	numerator := big.NewInt(1)
	denominator := big.NewInt(1)

	for other, share := range shares {
		if other == index {
			continue
		}

		numerator.Mul(numerator, s.difference(Origin, share.coordinate))
		denominator.Mul(denominator, s.difference(shares[index].coordinate, share.coordinate))
	}

	// Обратный элемент существует только для ненулевого знаменателя, то есть когда
	// координаты долей различны - иначе вернуть секрет невозможно.
	inverse := new(big.Int).ModInverse(denominator, s.prime)
	if inverse == nil {
		return nil, fmt.Errorf("coordinate %d: %w", shares[index].coordinate, ErrRepeatedCoordinate)
	}

	return numerator.Mul(numerator, inverse), nil
}

// difference - разность по модулю prime.
//
//	big.Int.Mod возвращает неотрицательный остаток, поэтому 0 - x_j считается
//	корректно, хотя x_j больше нуля.
func (s *Shamir) difference(left, right int64) *big.Int {
	difference := big.NewInt(left - right)

	return difference.Mod(difference, s.prime)
}

// validateSplit - проверки параметров до построения многочлена.
//
//	Число частей не может превышать размер поля, иначе координаты долей начнут
//	совпадать по модулю prime.
func (s *Shamir) validateSplit(secret int64, threshold, parts int) error {
	if threshold < MinShares {
		return fmt.Errorf("threshold %d: %w", threshold, ErrInvalidThreshold)
	}

	if parts < threshold {
		return fmt.Errorf("parts %d below threshold %d: %w", parts, threshold, ErrThresholdAboveParts)
	}

	if secret < Origin || secret >= s.prime.Int64() {
		return fmt.Errorf("secret %d: %w", secret, ErrSecretOutOfRange)
	}

	if parts >= int(s.prime.Int64()) {
		return fmt.Errorf("parts %d: %w", parts, ErrCoordinateOutOfRange)
	}

	return nil
}

// validateShares - проверки входных долей.
//
//	Координата должна лежать в 1..prime-1: при x = 0 значение многочлена это сам
//	секрет, а x = prime сравним с нулём по модулю поля. Значение должно лежать в
//	0..prime-1, иначе доли из разных полей нельзя смешивать.
func (s *Shamir) validateShares(shares []Share) error {
	if len(shares) < MinShares {
		return fmt.Errorf("got %d shares: %w", len(shares), ErrNotEnoughShares)
	}

	for _, share := range shares {
		if share.coordinate <= Origin || share.coordinate >= s.prime.Int64() {
			return fmt.Errorf("coordinate %d: %w", share.coordinate, ErrCoordinateOutOfRange)
		}

		if share.value == nil || share.value.Sign() < 0 || share.value.Cmp(s.prime) >= 0 {
			return fmt.Errorf("value %v: %w", share.value, ErrValueOutOfRange)
		}
	}

	return nil
}

// Демонстрация: делим секрет 13 на три доли с порогом 2, восстанавливаем его из
// двух и из трёх долей и показываем, что одна доля секрет не раскрывает.
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	shamir := NewShamir(DefaultPrime)

	shares, err := shamir.Split(DemoSecret, DemoThreshold, DemoParts)
	if err != nil {
		return fmt.Errorf("split secret: %w", err)
	}

	fmt.Printf(FormatSecret, DemoSecret)
	fmt.Printf(FormatShares, shares)

	// Порог достигнут и с двумя, и с тремя долями - восстановленный секрет одинаков.
	for _, subset := range [][]Share{shares[:DemoThreshold], shares} {
		recovered, err := shamir.Combine(subset)
		if err != nil {
			return fmt.Errorf("combine %d shares: %w", len(subset), err)
		}

		fmt.Printf(FormatRecovered, len(subset), recovered.Int64())
	}

	// Одна доля секрет не раскрывает: точек меньше, чем нужно для многочлена.
	_, err = shamir.Combine(shares[:MinShares-1])
	if !errors.Is(err, ErrNotEnoughShares) {
		return fmt.Errorf("combine single share: %w", err)
	}

	fmt.Printf(FormatBelowThreshold, MinShares-1)

	return nil
}
