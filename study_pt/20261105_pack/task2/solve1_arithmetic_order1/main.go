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
// ============================ ПОЧЕМУ НЕ HUFFMAN ===============================
//
// В тексте 26 разных букв.
// Если кодировать каждую букву отдельно, как в обычном Хаффмане, выйдет примерно 4.1-4.2 бита на букву.
// Это больше половины байта (4 бита), поэтому сжатие получится чуть хуже 50%.
// Значит, лимит "не больше половины" можно не пройти.
//
// Идея лучше: смотреть на предыдущую букву.
// В английском после 'q' почти всегда идёт 'u', а после 't' часто идёт 'h'.
// То есть вероятность буквы сильно зависит от контекста.
// Если кодировать букву с учётом предыдущей, в среднем выходит около 3.3-3.6 бита на букву.
// Это уже около 41-45%, поэтому запас есть.
//
// Такой код называют адаптивным арифметическим кодом порядка 1.
// "Арифметический" означает, что мы кодируем всю строку как одно длинное число-дробь в интервале [0, 1).
// При этом каждая буква сужает этот интервал.
// "Порядка 1" означает, что контекст имеет длину 1, то есть одна предыдущая буква.
// "Адаптивный" означает, что вероятности букв не хранятся в файле, а обновляются на ходу.
// После каждой буквы её частота увеличивается, и кодировщик с декодером делают это одинаково.
//
// ============================== КАК ЭТО РАБОТАЕТ ==============================
//
// Представим отрезок [0, 1).
// Делим его на 26 частей - по одной на каждую букву.
// Ширина части пропорциональна частоте буквы в текущем контексте.
// Сначала все частоты равны 1, то есть части одинаковые.
//
// Пример: алфавит из трёх букв a, b, c с частотами 2, 1, 1 (всего 4).
//   +--------+----+----+
//   |   a    | b  | c  |
//   +--------+----+----+
//   0      0.5  0.75  1.0
//
// Кодируем букву b - выбираем её отрезок [0.5, 0.75).
// Теперь работаем только внутри него и следующую букву делим уже внутри [0.5, 0.75).
// Чем больше букв, тем уже итоговый отрезок.
// В конце достаточно выдать любое число из этого отрезка, и декодер восстановит все буквы по одной.
//
// На бумаге это дроби, а в программе - целые числа low и high (границы отрезка) и арифметика на 32 битах.
// Когда low и high становятся слишком близки, старшие биты уже не изменятся.
// Их можно вывести в файл, и это и есть сжатие.
//
// ============================== ФОРМАТ ФАЙЛА =================================
//
// Сжатые данные - это просто массив байт:
//
//   байты 0..3          длина исходной строки n (little-endian, 4 байта)
//   байты 4..N-1        арифметически закодированные буквы
//
// Длина n нужна декодеру, чтобы знать, сколько букв разбирать.
// В арифметическом коде нет явного символа "конец строки", поэтому без длины непонятно, когда остановиться.
//
// ============================== СЛОЖНОСТЬ ====================================
//
// На каждую букву мы перебираем 26 частот, чтобы найти её границы.
// Итого O(26 * n) = O(n) по времени, потому что 26 - константа.
// Память: O(1) на модель (27 контекстов * 26 частот) плюс O(n) на результат.

const (
	codeBits     = 32 // сколько бит у границ интервала
	codeTop      = uint64(1) << codeBits
	codeHalf     = codeTop >> 1
	codeQuarter  = codeTop >> 2
	codeThreeQtr = codeQuarter * 3

	maxTotal = 1 << 15      // при достижении этой суммы частоты делим пополам
	alphabet = 26           // буквы 'a'..'z'
	numCtx   = alphabet + 1 // контекст 0 = начало строки, 1..26 = предыдущая буква
)

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
// Декодер при чтении за концом массива тоже считает недостающие биты нулями, поэтому всё сходится.
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
// Именно так буква "вписывается" в общий код.
func (e *encoder) encode(cumLow, cumHigh, total uint32) {
	r := e.high - e.low + 1 // ширина текущего интервала

	// Сначала считаем новый high по СТАРОМУ low, потом новый low.
	e.high = e.low + (r*uint64(cumHigh))/uint64(total) - 1
	e.low = e.low + (r*uint64(cumLow))/uint64(total)

	// Ренормализация: пока старшие биты границ совпадают или интервал попал в опасную середину, выводим биты и растягиваем интервал.
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
			// Обе границы в средней четверти: бит пока неизвестен, но следующий будет противоположен предыдущему.
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
// После него декодер сможет дочитать все буквы.
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

// scaled переводит текущее число value в номер внутри [0, total), по которому декодер ищет отрезок буквы.
func (d *decoder) scaled(total uint32) uint32 {
	r := d.high - d.low + 1
	return uint32(((d.value-d.low+1)*uint64(total) - 1) / r)
}

// decode делает то же сужение интервала, что и encoder.encode, но границы берутся из уже найденной буквы.
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
// После каждой буквы частота увеличивается на 1, поэтому модель подстраивается под текст.
// Декодер обновляет модель ровно так же, поэтому они не расходятся.
//
// Контекст ctx - это номер предыдущей буквы плюс 1.
// Значение ctx = 0 означает начало строки, где предыдущей буквы нет.
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
	ctx := 0 // начало строки: предыдущей буквы нет

	for i := 0; i < len(s); i++ {
		sym := int(s[i] - 'a') // 'a' -> 0, 'b' -> 1, ..., 'z' -> 25
		if sym < 0 || sym >= alphabet {
			sym = 0 // на всякий случай: во входе только a..z
		}

		lo, hi, total := m.rangeOf(ctx, sym)
		enc.encode(lo, hi, total)

		m.update(ctx, sym)
		ctx = sym + 1 // следующая буква кодируется с учётом текущей
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
	ctx := 0

	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
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
		ctx = sym + 1
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
