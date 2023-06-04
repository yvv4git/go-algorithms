package radix

import (
	"bytes"
	"encoding/binary"
)

func radixSortV1(list []int32) {
	/*
		Time complexity worst case: O(log_b(mx)(n+b))
		Time complexity average: O(d(n+b))
		Space complexity: O(n+b)

		Здесь,
		- b - основание системы счисления. У десятичной системы счисления это 10.
		- n - количество элементов в входном массиве.
		- d - количество цифр(разрядов) в числе.
	*/
	const wordLen = 4
	const highBit = -1 << 31

	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(list))
	for i, x := range list {
		binary.Write(buf, binary.LittleEndian, x^highBit)
		b := make([]byte, wordLen)
		buf.Read(b)
		ds[i] = b
	}

	bins := make([][][]byte, 256)
	for i := 0; i < wordLen; i++ {
		for _, b := range ds {
			bins[b[i]] = append(bins[b[i]], b)
		}
		j := 0
		for k, bs := range bins {
			copy(ds[j:], bs)
			j += len(bs)
			bins[k] = bs[:0]
		}
	}

	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		list[i] = w ^ highBit
	}
}
