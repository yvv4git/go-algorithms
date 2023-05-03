package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_powSHA256(t *testing.T) {
	type args struct {
		data       string
		difficulty int
	}

	type result struct {
		hash  []byte
		nonce int
	}

	tests := []struct {
		name string
		args args
		want result
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				data:       "Vladimir",
				difficulty: 1,
			},
			want: result{
				hash: []byte{
					0x0, 0xf8, 0x70, 0x1d, 0xa2, 0x36, 0x86, 0xf7, 0x5d, 0x56, 0xaf, 0xba, 0xbf, 0x97, 0x5, 0xca, 0x33,
					0x8d, 0x43, 0x5f, 0x1c, 0x81, 0x80, 0xf6, 0x3c, 0x23, 0x81, 0xcb, 0x8a, 0xec, 0x84, 0x7a,
				},
				nonce: 139,
			},
			desc: "Very fast ~ 0.00s",
		},
		{
			name: "CASE-2",
			args: args{
				data:       "Vladimir",
				difficulty: 2,
			},
			want: result{
				hash: []byte{
					0x0, 0x0, 0xbe, 0xc7, 0xf9, 0x66, 0x1, 0x77, 0x3b, 0xf3, 0xbf, 0x1e, 0xe3, 0x80, 0x32, 0xde, 0x53,
					0x4, 0x43, 0x9a, 0x9f, 0x95, 0xf3, 0x3b, 0x44, 0x5a, 0xe0, 0x59, 0xf9, 0x9d, 0x81, 0x3c,
				},
				nonce: 6952,
			},
			desc: "Also very fast ~ 0.00s",
		},
		{
			name: "CASE-3",
			args: args{
				data:       "Vladimir",
				difficulty: 3,
			},
			want: result{
				hash: []byte{
					0x0, 0x0, 0x0, 0x74, 0x9b, 0xb5, 0x11, 0x52, 0x6a, 0x51, 0x51, 0x42, 0x3, 0xd7, 0xc3, 0x1c, 0xe4,
					0x4e, 0xe7, 0x14, 0x9b, 0x31, 0xc1, 0xe5, 0x34, 0x70, 0x84, 0xf2, 0xa8, 0xdd, 0x43, 0xd1,
				},
				nonce: 12800919,
			},
			desc: "We have to wait ~5.43s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proof, nonce := powSHA256(tt.args.data, tt.args.difficulty)
			assert.Equal(t, tt.want.hash, proof)
			assert.Equal(t, tt.want.nonce, nonce)
			assert.True(t, checkProof([]byte(tt.args.data), proof, nonce))
		})
	}
}