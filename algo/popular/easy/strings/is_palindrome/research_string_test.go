package is_palindrome

import (
	"fmt"
	"testing"
	"unicode/utf8"
	"unsafe"
)

// -----
func TestResearch01(t *testing.T) {
	/*
		Каждый символ занимает int32.
	*/
	s := "вова"
	for _, c := range s {
		fmt.Printf("%c[%T %d] ", c, c, c)
	}
	fmt.Println()
}

func TestResearch02(t *testing.T) {
	/*
		Каждый символ занимает int32. В английском тоже.
	*/
	s := "vova"
	for _, c := range s {
		fmt.Printf("%c[%T %d] ", c, c, c)
	}
	fmt.Println()
}

func TestResearch03(t *testing.T) {
	/*
		Каждый символ занимает int32.
	*/
	s := "вова"
	sizeOfSymbol := unsafe.Sizeof(s[0])
	t.Logf("Size of symbol: %d", sizeOfSymbol) // 1 byte -> 8 bite -> int8
}

func TestResearch04(t *testing.T) {
	/*
		Каждый символ занимает int32.
	*/
	s := "вова"
	for _, c := range s {
		sizeOfSymbol := unsafe.Sizeof(c)
		t.Logf("Size of symbol: %d", sizeOfSymbol) // 4 byte -> 32 bit -> int32
	}
}

func TestResearch05(t *testing.T) {
	c := 'В'
	t.Logf("c=%c", c)
}

func TestResearch06(t *testing.T) {
	s := "вова"
	sLen := utf8.RuneCountInString(s)
	mid := sLen / 2
	var buf1 []rune
	t.Logf("mid=%d", mid)

	// Заполняем массив рун.
	// O(n) - на создание массива рун.
	for _, c := range s {
		fmt.Printf("%c ", c)
		buf1 = append(buf1, c)
	}
	fmt.Printf("\nbuf1=%#v \n", buf1)

	l, r := 0, len(buf1)-1
	for l < r {
		if buf1[l] != buf1[r] {
			t.Logf("l[%c] != r[%c]", buf1[l], buf1[r])
			break
		}
		l++
		r--
	}
	t.Log("+++")

	t.Log()
}

func TestResearch07(t *testing.T) {
	s := "вова"
	convertStrToRunesSlice := func(sInput string) []rune {
		var result []rune
		for _, r := range sInput {
			result = append(result, r)
		}

		return result
	}

	rList := convertStrToRunesSlice(s)

	t.Logf("String: %v", s)
	t.Logf("Runes list: %v", rList)

	isPalindromeStrRu := func(rl []rune) bool {
		l, r := 0, len(rList)-1
		for l < r {
			if rList[l] != rList[r] {
				return false
			}
			l++
			r--
		}

		return true
	}

	t.Logf("Result: %v", isPalindromeStrRu(rList))
}
