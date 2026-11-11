package main

import (
	"bufio"
	"fmt"
	"os"
)

// ============================== УСЛОВИЕ В ДВУХ СЛОВАХ =========================
//
// Программа запускается дважды.
//
// Первый запуск получает слово "pack" и строку s из маленьких английских букв (a..z).
// Нужно сжать строку и вывести две строки.
// В первой строке выводится N - сколько байт получилось.
// Во второй строке выводятся N чисел, каждое от 0 до 255.
//
// Второй запуск получает слово "unpack", затем N и N чисел из первого запуска.
// Нужно распаковать данные и вывести исходную строку.
//
// Главное требование: после распаковки получается ровно исходная строка.
// Размер сжатых данных не должен превышать половины длины исходной строки.
//
// ==================== ЧЕМ ЭТО ОТЛИЧАЕТСЯ ОТ solve1 ============================
//
// В solve1 использовался контекст из одной предыдущей буквы (порядок 1).
// Здесь используется контекст из двух предыдущих букв (порядок 2).
//
// В английском пары "th", "he", "in", "er" встречаются очень часто.
// Они предсказывают следующую букву лучше, чем одна буква.
// Поэтому сжатие получается плотнее: около 2.9-3.2 бита на букву вместо 3.3-3.6.
//
// ============================ ПОЧЕМУ НЕ HUFFMAN ===============================
//
// Обычный Хаффман по одной букве даёт около 4.1-4.2 бита на букву.
// Это больше половины байта, и лимит "не больше половины" можно не пройти.
// Контекст из двух букв снижает неопределённость, и результат уверенно укладывается в 50%.
//
// ============================== КАК ЭТО РАБОТАЕТ ==============================
//
// Арифметическое кодирование представляет весь текст как одно число-дробь из отрезка [0, 1).
// Каждая буква сужает текущий отрезок.
//
// Пример: алфавит из трёх букв a, b, c с частотами 2, 1, 1 (всего 4).
//   +--------+----+----+
//   |   a    | b  | c  |
//   +--------+----+----+
//   0      0.5  0.75  1.0
//
// Кодируем b - переходим к отрезку [0.5, 0.75).
// Следующую букву делим уже внутри него, и так далее.
// Чем больше букв, тем уже отрезок.
// В конце выдаём любое число из отрезка, и декодер по нему восстанавливает все буквы.
//
// В программе вместо дробей - целые low и high на 32 битах.
// Когда старшие биты low и high совпадают, они уже не изменятся, и их выводят в файл.
//
// Вероятности берутся из модели.
// Модель адаптивная: после каждой буквы её частота увеличивается.
// Кодировщик и декодер обновляют модель одинаково, поэтому не расходятся.
//
// ============================== КОНТЕКСТЫ ====================================
//
// Контекст - это то, что мы знаем о предыдущих буквах:
//
//   context 0          - начало строки (предыдущих букв нет);
//   context 1..26      - известна одна предыдущая буква (A);
//   context 27..702    - известны две предыдущие буквы (A, B).
//
// Для двух букв номер контекста считается так:
//
//   ctx = 27 + A * 26 + B
//
// где A и B - номера букв от 0 до 25.
//
// Модель частот - массив 703 x 26: для каждого контекста своя таблица частот.
// Это меньше 100 КБ, поэтому память не проблема.
//
// ============================== ФОРМАТ ФАЙЛА =================================
//
//   байты 0..3          длина исходной строки n (little-endian, 4 байта)
//   байты 4..N-1        арифметически закодированные буквы
//
// Длина нужна декодеру, чтобы знать, сколько букв разбирать.
// В арифметическом коде нет явного символа "конец строки".
//
// ============================== СЛОЖНОСТЬ ====================================
//
// На каждую букву перебираем 26 частот, поэтому O(26 * n) = O(n) по времени.
// Память: O(1) на модель (фиксированные 703 * 26) плюс O(n) на результат.

const (
	codeBits     = 32 // сколько бит у границ интервала
	codeTop      = uint64(1) << codeBits
	codeHalf     = codeTop >> 1
	codeQuarter  = codeTop >> 2
	codeThreeQtr = codeQuarter * 3

	maxTotal = 1 << 15 // при достижении этой суммы частоты делим пополам
	alphabet = 26      // буквы 'a'..'z'

	// Контексты трёх видов: начало строки, одна буква, две буквы.
	numCtx = 1 + alphabet + alphabet*alphabet
)

// contextOf выбирает номер контекста по двум предыдущим буквам.
// prev2 и prev1 равны -1, если соответствующей буквы ещё нет.
func contextOf(prev2, prev1 int) int {
	if prev1 < 0 {
		return 0 // начало строки
	}
	if prev2 < 0 {
		return 1 + prev1 // известна только одна предыдущая буква
	}
	return 1 + alphabet + prev2*alphabet + prev1 // известны две предыдущие буквы
}

// ============================== ВЫВОД БИТ ====================================
//
// Арифметический код выдаёт отдельные биты, а нам нужны байты.
// bitWriter складывает биты в байт, старший бит первым (MSB-first).
type bitWriter struct {
	buf   []byte
	cur   byte
	nbits int
}

func (w *bitWriter) writeBit(bit int) {
	w.cur = (w.cur << 1) | byte(bit)
	w.nbits++
	if w.nbits == 8 {
		w.buf = append(w.buf, w.cur)
		w.cur = 0
		w.nbits = 0
	}
}

// flush дописывает последний неполный байт.
// Для этого он сдвигает байт влево и заполняет младшие биты нулями.
// Декодер при чтении за концом массива тоже считает недостающие биты нулями.
func (w *bitWriter) flush() {
	if w.nbits > 0 {
		w.cur <<= (8 - w.nbits)
		w.buf = append(w.buf, w.cur)
		w.cur = 0
		w.nbits = 0
	}
}

// ============================== КОДИРОВЩИК ===================================
type encoder struct {
	low     uint64 // нижняя граница текущего интервала
	high    uint64 // верхняя граница текущего интервала
	pending int    // отложенные биты для правила E3
	bw      *bitWriter
}

func newEncoder() *encoder {
	return &encoder{low: 0, high: codeTop - 1, bw: &bitWriter{}}
}

// encode сужает интервал [low, high) до части [cumLow, cumHigh) из total.
func (e *encoder) encode(cumLow, cumHigh, total uint32) {
	r := e.high - e.low + 1 // ширина текущего интервала

	// Сначала считаем новый high по СТАРОМУ low, потом новый low.
	e.high = e.low + (r*uint64(cumHigh))/uint64(total) - 1
	e.low = e.low + (r*uint64(cumLow))/uint64(total)

	// Ренормализация: выводим биты, которые уже не изменятся.
	for {
		if e.high < codeHalf {
			// Обе границы в нижней половине: следующий бит точно 0.
			e.outputBit(0)
		} else if e.low >= codeHalf {
			// Обе границы в верхней половине: следующий бит точно 1.
			e.outputBit(1)
			e.low -= codeHalf
			e.high -= codeHalf
		} else if e.low >= codeQuarter && e.high < codeThreeQtr {
			// Средняя четверть: бит пока неизвестен, откладываем его.
			e.pending++
			e.low -= codeQuarter
			e.high -= codeQuarter
		} else {
			break
		}
		e.low <<= 1
		e.high = e.high<<1 | 1
	}
}

// outputBit выводит бит и за ним все отложенные биты в противоположном значении.
// Это стандартное правило E3 арифметического кодирования.
func (e *encoder) outputBit(bit int) {
	e.bw.writeBit(bit)
	for ; e.pending > 0; e.pending-- {
		e.bw.writeBit(1 - bit)
	}
}

// finish дописывает хвост кода.
func (e *encoder) finish() {
	e.pending++
	if e.low < codeQuarter {
		e.outputBit(0)
	} else {
		e.outputBit(1)
	}
	e.bw.flush()
}

// ============================== ЧТЕНИЕ БИТ ===================================
type bitReader struct {
	data  []byte
	idx   int
	cur   byte
	nbits int
}

// readBit читает очередной бит, старший в байте.
// Если данные закончились, возвращает 0, потому что кодировщик добивает хвост нулями.
func (r *bitReader) readBit() int {
	if r.nbits == 0 {
		if r.idx >= len(r.data) {
			return 0
		}
		r.cur = r.data[r.idx]
		r.idx++
		r.nbits = 8
	}
	r.nbits--
	return int((r.cur >> uint(r.nbits)) & 1)
}

// ============================== ДЕКОДЕР ======================================
type decoder struct {
	low   uint64
	high  uint64
	value uint64 // число, которое "накопилось" из прочитанных битов
	br    *bitReader
}

func newDecoder(data []byte) *decoder {
	d := &decoder{low: 0, high: codeTop - 1, br: &bitReader{data: data}}
	for i := 0; i < codeBits; i++ {
		d.value = d.value<<1 | uint64(d.br.readBit())
	}
	return d
}

// scaled переводит текущее число value в номер внутри [0, total).
func (d *decoder) scaled(total uint32) uint32 {
	r := d.high - d.low + 1
	return uint32(((d.value-d.low+1)*uint64(total) - 1) / r)
}

// decode делает то же сужение интервала, что и encoder.encode.
func (d *decoder) decode(cumLow, cumHigh, total uint32) {
	r := d.high - d.low + 1
	d.high = d.low + (r*uint64(cumHigh))/uint64(total) - 1
	d.low = d.low + (r*uint64(cumLow))/uint64(total)

	for {
		if d.high < codeHalf {
			// нижняя половина - ничего не вычитаем
		} else if d.low >= codeHalf {
			d.low -= codeHalf
			d.high -= codeHalf
			d.value -= codeHalf
		} else if d.low >= codeQuarter && d.high < codeThreeQtr {
			d.low -= codeQuarter
			d.high -= codeQuarter
			d.value -= codeQuarter
		} else {
			break
		}
		d.low <<= 1
		d.high = d.high<<1 | 1
		d.value = d.value<<1 | uint64(d.br.readBit())
	}
}

// ============================== МОДЕЛЬ ЧАСТОТ ================================
//
// freq[ctx][sym] - сколько раз символ sym встретился после контекста ctx.
// Изначально у всех частота 1, то есть все буквы равновероятны.
// После каждой буквы её частота увеличивается.
// Декодер обновляет модель так же, как кодировщик.
type model struct {
	freq  [numCtx][alphabet]uint32
	total [numCtx]uint32
}

func newModel() *model {
	m := &model{}
	for ctx := 0; ctx < numCtx; ctx++ {
		for sym := 0; sym < alphabet; sym++ {
			m.freq[ctx][sym] = 1
		}
		m.total[ctx] = alphabet
	}
	return m
}

// rangeOf возвращает границы [lo, hi) символа sym в контексте ctx и сумму всех частот контекста.
func (m *model) rangeOf(ctx, sym int) (lo, hi, total uint32) {
	for i := 0; i < sym; i++ {
		lo += m.freq[ctx][i]
	}
	hi = lo + m.freq[ctx][sym]
	return lo, hi, m.total[ctx]
}

// update увеличивает частоту символа.
// Если сумма частот стала слишком большой, все частоты уменьшаются вдвое.
// Иначе арифметике не хватит точности.
// Кодировщик и декодер делают это синхронно.
func (m *model) update(ctx, sym int) {
	m.freq[ctx][sym]++
	m.total[ctx]++

	if m.total[ctx] >= maxTotal {
		var sum uint32
		for i := 0; i < alphabet; i++ {
			m.freq[ctx][i] = (m.freq[ctx][i] + 1) / 2
			sum += m.freq[ctx][i]
		}
		m.total[ctx] = sum
	}
}

// ============================== PACK / UNPACK ================================

// pack сжимает строку s и возвращает массив байт.
//
// Первые 4 байта - длина исходной строки, дальше - арифметический код.
func pack(s string) []byte {
	n := uint32(len(s))
	out := make([]byte, 0, len(s)/2+16)
	out = append(out, byte(n), byte(n>>8), byte(n>>16), byte(n>>24))

	enc := newEncoder()
	m := newModel()
	prev2, prev1 := -1, -1 // двух предыдущих букв ещё нет

	for i := 0; i < len(s); i++ {
		sym := int(s[i] - 'a') // 'a' -> 0, 'b' -> 1, ..., 'z' -> 25
		if sym < 0 || sym >= alphabet {
			sym = 0 // на всякий случай: во входе только a..z
		}

		ctx := contextOf(prev2, prev1)
		lo, hi, total := m.rangeOf(ctx, sym)
		enc.encode(lo, hi, total)

		m.update(ctx, sym)
		prev2, prev1 = prev1, sym
	}

	enc.finish()
	out = append(out, enc.bw.buf...)
	return out
}

// unpack восстанавливает строку из массива байт, полученного pack.
func unpack(data []byte) string {
	if len(data) < 4 {
		return ""
	}

	// Читаем длину исходной строки.
	n := int(uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24)

	dec := newDecoder(data[4:])
	m := newModel()
	prev2, prev1 := -1, -1

	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		ctx := contextOf(prev2, prev1)
		total := m.total[ctx]
		target := dec.scaled(total) // номер, в который попала текущая буква

		// Ищем, какому символу соответствует этот номер.
		var cum uint32
		sym := 0
		for ; sym < alphabet; sym++ {
			if cum+m.freq[ctx][sym] > target {
				break
			}
			cum += m.freq[ctx][sym]
		}
		if sym == alphabet {
			sym = alphabet - 1 // защита от выхода за границу
		}

		lo := cum
		hi := cum + m.freq[ctx][sym]
		dec.decode(lo, hi, total)

		result = append(result, byte('a'+sym))
		m.update(ctx, sym)
		prev2, prev1 = prev1, sym
	}
	return string(result)
}

// ============================== MAIN =========================================
//
// Читает режим и данные, вызывает pack или unpack.
// Формат вывода строго по условию: сначала N, потом N чисел через пробел.
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var command string
	if _, err := fmt.Fscan(in, &command); err != nil {
		return
	}

	switch command {
	case "pack":
		var s string
		if _, err := fmt.Fscan(in, &s); err != nil {
			return
		}

		data := pack(s)
		fmt.Fprintln(out, len(data))
		for i, b := range data {
			if i > 0 {
				out.WriteByte(' ')
			}
			fmt.Fprint(out, b)
		}
		fmt.Fprintln(out)

	case "unpack":
		var n int
		if _, err := fmt.Fscan(in, &n); err != nil {
			return
		}

		data := make([]byte, n)
		for i := 0; i < n; i++ {
			var v int
			if _, err := fmt.Fscan(in, &v); err != nil {
				return
			}
			data[i] = byte(v)
		}

		fmt.Fprintln(out, unpack(data))
	}
}
