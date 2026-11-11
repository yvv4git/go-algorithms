package main

import (
	"strings"
	"testing"
)

func Test_packUnpack(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"single letter", "a"},
		{"last letter", "z"},
		{"example from statement", "abacabaca"},
		{"word", "hello"},
		{"pangram", "thequickbrownfoxjumpsoverthelazydog"},
		{"periodic", strings.Repeat("abcabcabc", 1000)},
		{"all letters", strings.Repeat("abcdefghijklmnopqrstuvwxyz", 500)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unpack(pack(tt.input)); got != tt.input {
				t.Fatalf("round trip failed: got %q, want %q", got, tt.input)
			}
		})
	}
}

func Test_packUnpack_maxSize(t *testing.T) {
	text := EnglishSample(200000)
	data := pack(text)

	if got := unpack(data); got != text {
		t.Fatalf("round trip failed on maximum size text")
	}
	if len(data)*2 > len(text) {
		t.Fatalf("compressed size %d is more than half of %d", len(data), len(text))
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
