package main

import (
	"reflect"
	"testing"
)

func Test_quickselect(t *testing.T) {
	type args struct {
		results []SearchResult
		k       int
	}

	tests := []struct {
		name string
		args args
		want SearchResult
	}{
		{
			name: "Case 1: k=1, smallest element",
			args: args{
				results: []SearchResult{
					{ID: 1, Relevance: 5},
					{ID: 2, Relevance: 3},
					{ID: 3, Relevance: 8},
					{ID: 4, Relevance: 2},
					{ID: 5, Relevance: 7},
				},
				k: 1,
			},
			want: SearchResult{ID: 4, Relevance: 2},
		},
		{
			name: "Case 2: k=3, middle element",
			args: args{
				results: []SearchResult{
					{ID: 1, Relevance: 5},
					{ID: 2, Relevance: 3},
					{ID: 3, Relevance: 8},
					{ID: 4, Relevance: 2},
					{ID: 5, Relevance: 7},
				},
				k: 3,
			},
			want: SearchResult{ID: 1, Relevance: 5},
		},
		{
			name: "Case 3: k=5, largest element",
			args: args{
				results: []SearchResult{
					{ID: 1, Relevance: 5},
					{ID: 2, Relevance: 3},
					{ID: 3, Relevance: 8},
					{ID: 4, Relevance: 2},
					{ID: 5, Relevance: 7},
				},
				k: 5,
			},
			want: SearchResult{ID: 3, Relevance: 8},
		},
		{
			name: "Case 4: k=1, single element array",
			args: args{
				results: []SearchResult{
					{ID: 1, Relevance: 5},
				},
				k: 1,
			},
			want: SearchResult{ID: 1, Relevance: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickselect(tt.args.results, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickselect() = %v, want %v", got, tt.want)
			}
		})
	}
}
