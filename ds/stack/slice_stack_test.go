package stack

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

func Test_initStackInt64(t *testing.T) {
	type args struct {
		values    []int64
		batchSize int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "CASE-0",
			args: args{
				values:    []int64{},
				batchSize: 5,
			},
		},
		{
			name: "CASE-1",
			args: args{
				values:    genSequence(10),
				batchSize: 3,
			},
		},
		{
			name: "CASE-2",
			args: args{
				values:    genSequence(100),
				batchSize: 5,
			},
		},
		{
			name: "CASE-3",
			args: args{
				values:    genSequence(100),
				batchSize: 15,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pop := initStackInt64(tt.args.values)

			for {
				result, err := pop(tt.args.batchSize)
				if err != nil {
					if !errors.Is(err, ErrGoingOffStack) {
						t.Fatalf("problems with stack: %v", err)
					}

					break
				}

				t.Log(result)
			}
		})
	}
}
