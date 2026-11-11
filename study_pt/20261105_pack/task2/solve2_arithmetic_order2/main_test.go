package main

import (
	"strings"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	texts := []string{
		"",
		"a",
		"z",
		"hello",
		"thequickbrownfoxjumpsoverthelazydog",
		strings.Repeat("abcabcabc", 1000),
	}

	for _, text := range texts {
		got := unpack(pack(text))
		if got != text {
			t.Fatalf("round trip failed:\n in = %q\nout = %q", text, got)
		}
	}
}

func EnglishSample(n int) string {
	paragraph := "itwasabrightcolddayinaprilandtheclockswerestrikingthirteenwinstonsmithischinnuzzledintohis" +
		"breastinanefforttoescapethevilewindslippedquicklythroughtheglassdoorsofvictorymansionsthoughnot" +
		"quicklyenoughtopreventaswirlofgrittydustfromenteringalongwithhimthehallwaysmeltsofboiledcabbage" +
		"andoldragmatsatoneendofitacolouredpostertoolargeforindoordisplayhadbeentackedtothewallitdepicted" +
		"simplyanenormousfacemorethanametrewidethefaceofamanofaboutfortyfivewithaheavyblackmoustacheand" +
		"ruggedlyhandsomefeatures"
	var b strings.Builder
	for b.Len() < n {
		b.WriteString(paragraph)
	}
	return b.String()[:n]
}

func TestCompressionRatio(t *testing.T) {
	const n = 100000
	text := EnglishSample(n)
	data := pack(text)

	if got := unpack(data); got != text {
		t.Fatalf("round trip failed on long text")
	}
	if len(data)*2 > len(text) {
		t.Fatalf("compressed size %d is more than half of %d", len(data), len(text))
	}
	t.Logf("input=%d bytes, compressed=%d bytes, ratio=%.2f%%",
		len(text), len(data), float64(len(data))*100/float64(len(text)))
}

func TestAllLetterPairs(t *testing.T) {
	var b strings.Builder
	for i := 0; i < 500; i++ {
		b.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	text := b.String()
	if got := unpack(pack(text)); got != text {
		t.Fatalf("round trip on alphabet failed")
	}
}
