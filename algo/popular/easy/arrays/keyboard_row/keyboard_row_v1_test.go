package keyboard_row

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "CASE-1",
			args: args{
				words: []string{"Hello", "Alaska", "Dad", "Peace"},
			},
			want: []string{"Alaska", "Dad"},
		},
		{
			name: "CASE-2",
			args: args{
				words: []string{"omk"},
			},
			want: []string{},
		},
		{
			name: "CASE-3",
			args: args{
				words: []string{"adsdf", "sfd"},
			},
			want: []string{"adsdf", "sfd"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordsV1(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------------------------------------------------------------------
func TestResearch01(t *testing.T) {
	for _, num := range []byte("qwertyuiop") {
		v := int(num - 'a')
		t.Logf("v=%v", v)
	}
}
