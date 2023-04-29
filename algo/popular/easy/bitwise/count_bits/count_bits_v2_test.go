package count_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cntOnes(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "CASE-2",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "CASE-3",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "CASE-4",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "CASE-5",
			args: args{
				n: 4,
			},
			want: 1,
		},
		{
			name: "CASE-6",
			args: args{
				n: 5,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cntOnes(tt.args.n)
			t.Logf("ones count in number: %d [n=%d as %08b]", result, tt.args.n, tt.args.n)
			assert.Equal(t, tt.want, result)
		})
	}
}

// ---------------------------------------------
func TestResearch_calcCountOnes(t *testing.T) {
	/*
		Здесь я просто визуализирую для себя, как работает эта операция.
	*/
	n := 5
	cnt := 0
	for n > 0 {
		t.Logf(".. n=%d(%08b) cnt=%d n&1=%d", n, n, cnt, n&1)
		cnt += n & 1
		n >>= 1
		t.Logf("-> n=%d(%08b) cnt=%d", n, n, cnt)
	}

	t.Logf("In the end n=%d(%08b) cnt=%d", n, n, cnt)
}
