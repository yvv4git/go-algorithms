package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// Blakley Secret Sharing (1979) - геометрическая схема разделения секрета,
// вторая классическая работа наряду с Shamir.
//
//	Секрет - первая координата точки T = (S, a_1, ..., a_{M-1}) в поле.
//	Доля i - гиперплоскость, проходящая через эту точку:
//
//	y_i = S + b_{i1}·a_1 + b_{i2}·a_2 + ... + b_{i,M-1}·a_{M-1}
//
//	Восстановление - пересечение M гиперплоскостей, то есть решение системы
//	из M линейных уравнений по отношению к неизвестным (x_0 = S, x_1, ...).
//
//	В отличие от Shamir, доля хранит M чисел вместо одного, а восстановление
//	требует решения СЛАУ методом Гаусса.
const (
	DefaultPrime int64 = 257 // размер поля, все числа многочлена берутся по модулю prime
	MinShares    int   = 2   // меньше двух уравнений точку не определяют
	SecretColumn int   = 0   // номер неизвестного, которым является секрет
)

// Параметры демонстрации: секрет 13 делится на 3 доли с порогом 2,
// поэтому каждая доля - это одно уравнение с двумя неизвестными.
const (
	DemoSecret    int64 = 13
	DemoThreshold int   = 2
	DemoParts     int   = 3
)

// Форматы вывода вынесены в константы, чтобы vet мог проверять аргументы Printf.
const (
	FormatSecret         = "secret: %d\n"
	FormatShareLine      = "  share: %v\n"
	FormatPlaneSecret    = "y = x0"
	FormatPlaneTerm      = " + %d·x%d"
	FormatPlaneValue     = " = %d"
	FormatRecovered      = "recovered from %d planes: %d\n"
	FormatBelowThreshold = "%d plane below the threshold cannot fix the point\n"
)

var (
	ErrThresholdBelowMin   = errors.New("threshold needs at least two shares")
	ErrPartsBelowThreshold = errors.New("parts below threshold")
	ErrSecretOutOfRange    = errors.New("secret outside the field")
	ErrNotEnoughShares     = errors.New("not enough planes for the unknowns")
	ErrShareArity          = errors.New("plane with unexpected number of coefficients")
	ErrValueOutOfRange     = errors.New("value outside the field")
	ErrSingularSystem      = errors.New("hyperplanes do not intersect in a single point")
)

// Share - доля Blakley: коэффициенты гиперплоскости b_1..b_{M-1} и её y.
//
//	Доля это целое уравнение, поэтому её размер больше, чем у доли Shamir.
type Share struct {
	coefficients []*big.Int
	value        *big.Int
}

// String - доля как уравнение гиперплоскости, например "y = x0 + 3·x1 = 19".
func (share Share) String() string {
	plane := new(strings.Builder)
	plane.WriteString(FormatPlaneSecret)

	for index, coefficient := range share.coefficients {
		fmt.Fprintf(plane, FormatPlaneTerm, coefficient, index+1)
	}

	fmt.Fprintf(plane, FormatPlaneValue, share.value)

	return plane.String()
}

// Blakley - разделение секрета гиперплоскостями над полем по модулю prime.
type Blakley struct {
	prime *big.Int
	// draw - источник случайных коэффициентов, поле для подмены в тестах.
	draw func(limit *big.Int) (*big.Int, error)
}

func NewBlakley(prime int64) *Blakley {
	draw := func(limit *big.Int) (*big.Int, error) { return rand.Int(rand.Reader, limit) }

	return &Blakley{prime: big.NewInt(prime), draw: draw}
}

// Split - разбиение секрета на parts гиперплоскостей с порогом threshold.
//
//	METHOD: выбирается случайная точка T = (S, a_1, ..., a_{M-1}), затем для
//	каждой доли случайные b_1..b_{M-1} и y = S + Σ b_j·a_j (mod p).
//
//	Любые M долей дают M уравнений на M неизвестных, а значит единственную точку,
//	первая координата которой и есть секрет. M-1 долей оставляют M-1 свободных
//	координат точки, поэтому секрет не определяется.
//
//	TIME COMPLEXITY: O(parts * threshold).
func (b *Blakley) Split(secret int64, threshold, parts int) ([]Share, error) {
	if err := b.validateSplit(secret, threshold, parts); err != nil {
		return nil, err
	}

	point, err := b.coordinates(threshold - 1)
	if err != nil {
		return nil, fmt.Errorf("build point: %w", err)
	}

	shares := make([]Share, 0, parts)
	for range parts {
		share, err := b.plane(secret, point)
		if err != nil {
			return nil, fmt.Errorf("build plane: %w", err)
		}

		shares = append(shares, share)
	}

	return shares, nil
}

// Combine - восстановление секрета как пересечение гиперплоскостей.
//
//	METHOD: система из M уравнений
//
//		y_i = x_0 + b_{i1}·x_1 + ... + b_{i,M-1}·x_{M-1}
//
//	решается методом Гаусса по модулю prime, секрет - это x_0.
//
//	TIME COMPLEXITY: O(threshold^3), против O(threshold^2) у Shamir.
func (b *Blakley) Combine(shares []Share) (*big.Int, error) {
	if err := b.validateShares(shares); err != nil {
		return nil, err
	}

	system, err := b.system(shares)
	if err != nil {
		return nil, err
	}

	solution, err := b.solve(system)
	if err != nil {
		return nil, fmt.Errorf("solve system: %w", err)
	}

	return solution[SecretColumn], nil
}

// coordinates - count случайных координат из поля, координаты точки или плоскости.
func (b *Blakley) coordinates(count int) ([]*big.Int, error) {
	coordinates := make([]*big.Int, 0, count)

	for range count {
		coordinate, err := b.draw(b.prime)
		if err != nil {
			return nil, fmt.Errorf("draw coordinate: %w", err)
		}

		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

// plane - случайная гиперплоскость через точку: y = S + Σ b_j · a_j.
//
//	y считается подстановкой, поэтому плоскость гарантированно проходит через
//	точку, а значит любые threshold плоскостей пересекаются в этой точке.
func (b *Blakley) plane(secret int64, point []*big.Int) (Share, error) {
	coefficients, err := b.coordinates(len(point))
	if err != nil {
		return Share{}, err
	}

	value := big.NewInt(secret)
	for index, coordinate := range point {
		term := new(big.Int).Mul(coefficients[index], coordinate)
		value.Add(value, term).Mod(value, b.prime)
	}

	return Share{coefficients: coefficients, value: value}, nil
}

// system - расширенная матрица из threshold уравнений, готовая к решению.
//
//	Число неизвестных определяется по арности первой доли: M = коэффициентов + 1.
func (b *Blakley) system(shares []Share) ([][]*big.Int, error) {
	unknowns := len(shares[0].coefficients) + 1
	if len(shares) < unknowns {
		return nil, fmt.Errorf("%d planes for %d unknowns: %w", len(shares), unknowns, ErrNotEnoughShares)
	}

	system := make([][]*big.Int, 0, unknowns)
	for _, share := range shares[:unknowns] {
		row, err := b.row(share, unknowns)
		if err != nil {
			return nil, err
		}

		system = append(system, row)
	}

	return system, nil
}

// row - строка расширенной матрицы: 1 при секрете, коэффициенты плоскости и y.
//
//	Значения копируются, потому что метод Гаусса меняет матрицу на месте.
func (b *Blakley) row(share Share, unknowns int) ([]*big.Int, error) {
	if len(share.coefficients) != unknowns-1 {
		return nil, fmt.Errorf("plane with %d coefficients: %w", len(share.coefficients), ErrShareArity)
	}

	row := make([]*big.Int, 0, unknowns+1)
	row = append(row, big.NewInt(1))

	for _, coefficient := range share.coefficients {
		if err := b.validateValue(coefficient); err != nil {
			return nil, err
		}

		row = append(row, new(big.Int).Set(coefficient))
	}

	return append(row, new(big.Int).Set(share.value)), nil
}

// solve - метод Гаусса с выбором ведущего элемента по первому ненулевому.
//
//	После приведения к единичной диагонали ответ лежит в последнем столбце.
func (b *Blakley) solve(system [][]*big.Int) ([]*big.Int, error) {
	size := len(system)

	for column := range size {
		pivot, err := b.pivotIndex(system, column)
		if err != nil {
			return nil, err
		}

		system[column], system[pivot] = system[pivot], system[column]
		b.normalize(system, column)

		for row := range size {
			if row == column {
				continue
			}

			b.eliminate(system, column, row)
		}
	}

	solution := make([]*big.Int, 0, size)
	for _, row := range system {
		solution = append(solution, row[size])
	}

	return solution, nil
}

// pivotIndex - номер строки с ненулевым элементом в столбце, начиная с текущей.
//
//	Ненулевого элемента нет - значит система вырождена и точку не определить.
func (b *Blakley) pivotIndex(system [][]*big.Int, column int) (int, error) {
	for index := column; index < len(system); index++ {
		if system[index][column].Sign() != 0 {
			return index, nil
		}
	}

	return 0, fmt.Errorf("column %d: %w", column, ErrSingularSystem)
}

// normalize - делит ведущую строку на ведущий элемент, делая его равным 1.
func (b *Blakley) normalize(system [][]*big.Int, pivot int) {
	inverse := new(big.Int).ModInverse(system[pivot][pivot], b.prime)
	lastColumn := len(system[pivot]) - 1

	for column := pivot; column <= lastColumn; column++ {
		system[pivot][column].Mul(system[pivot][column], inverse).Mod(system[pivot][column], b.prime)
	}
}

// eliminate - обнуляет элемент столбца в строке row вычитанием ведущей строки.
func (b *Blakley) eliminate(system [][]*big.Int, pivot, row int) {
	factor := new(big.Int).Set(system[row][pivot])
	if factor.Sign() == 0 {
		return
	}

	lastColumn := len(system[row]) - 1
	for column := pivot; column <= lastColumn; column++ {
		product := new(big.Int).Mul(factor, system[pivot][column])
		system[row][column].Sub(system[row][column], product).Mod(system[row][column], b.prime)
	}
}

func (b *Blakley) validateSplit(secret int64, threshold, parts int) error {
	if threshold < MinShares {
		return fmt.Errorf("threshold %d: %w", threshold, ErrThresholdBelowMin)
	}

	if parts < threshold {
		return fmt.Errorf("parts %d below threshold %d: %w", parts, threshold, ErrPartsBelowThreshold)
	}

	if secret < 0 || secret >= b.prime.Int64() {
		return fmt.Errorf("secret %d: %w", secret, ErrSecretOutOfRange)
	}

	return nil
}

func (b *Blakley) validateShares(shares []Share) error {
	if len(shares) < MinShares {
		return fmt.Errorf("got %d planes: %w", len(shares), ErrNotEnoughShares)
	}

	for _, share := range shares {
		if len(share.coefficients) < MinShares-1 {
			return fmt.Errorf("plane with %d coefficients: %w", len(share.coefficients), ErrShareArity)
		}

		if err := b.validateValue(share.value); err != nil {
			return err
		}
	}

	return nil
}

func (b *Blakley) validateValue(value *big.Int) error {
	if value == nil || value.Sign() < 0 || value.Cmp(b.prime) >= 0 {
		return fmt.Errorf("value %v: %w", value, ErrValueOutOfRange)
	}

	return nil
}

// Демонстрация: секрет 13 прячется в точке за тремя гиперплоскостями, любые две
//
//	из них пересекаются в этой точке, а одна плоскость точку не определяет.
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	blakley := NewBlakley(DefaultPrime)

	shares, err := blakley.Split(DemoSecret, DemoThreshold, DemoParts)
	if err != nil {
		return fmt.Errorf("split secret: %w", err)
	}

	fmt.Printf(FormatSecret, DemoSecret)
	for _, share := range shares {
		fmt.Printf(FormatShareLine, share)
	}

	secret, err := blakley.Combine(shares[:DemoThreshold])
	if err != nil {
		return fmt.Errorf("combine shares: %w", err)
	}

	fmt.Printf(FormatRecovered, DemoThreshold, secret.Int64())

	_, err = blakley.Combine(shares[:DemoThreshold-1])
	if !errors.Is(err, ErrNotEnoughShares) {
		return fmt.Errorf("combine below threshold: %w", err)
	}

	fmt.Printf(FormatBelowThreshold, DemoThreshold-1)

	return nil
}
