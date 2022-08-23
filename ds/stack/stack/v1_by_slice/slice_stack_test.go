package v1_by_slice

import (
	"errors"
	"testing"
)

func genSequence(count int64) []int64 {
	if count == 0 {
		return []int64{}
	}

	result := []int64{}
	for i := 1; i <= int(count); i++ {
		result = append(result, int64(i))
	}

	return result
}

func Test_NewStackInt64(t *testing.T) {
	type args struct {
		values    []int64
		batchSize int
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "CASE-0",
			args: args{
				values:    []int64{},
				batchSize: 5,
			},
			err: ErrEmptyStack,
		},
		{
			name: "CASE-1",
			args: args{
				values:    genSequence(10),
				batchSize: 3,
			},
			err: ErrGoingOffStack,
		},
		{
			name: "CASE-2",
			args: args{
				values:    genSequence(100),
				batchSize: 5,
			},
			err: ErrGoingOffStack,
		},
		{
			name: "CASE-3",
			args: args{
				values:    genSequence(100),
				batchSize: 15,
			},
			err: ErrGoingOffStack,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pop := NewStackInt64(tt.args.values)

			for {
				_, err := pop(tt.args.batchSize)
				if err != nil {
					if !errors.Is(err, tt.err) {
						t.Fatalf("we want: %v, but got %v", tt.err, err)
					}

					break
				}
			}
		})
	}
}
