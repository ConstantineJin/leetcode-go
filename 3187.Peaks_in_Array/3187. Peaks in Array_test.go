package main

import (
	"reflect"
	"testing"
)

func Test_countOfPeaks(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				nums:    []int{3, 1, 4, 2, 5},
				queries: [][]int{{2, 3, 4}, {1, 0, 4}},
			},
			wantAns: []int{0},
		},
		{
			name: "Example2",
			args: args{
				nums:    []int{4, 1, 4, 2, 1, 5},
				queries: [][]int{{2, 2, 4}, {1, 0, 2}, {1, 0, 4}},
			},
			wantAns: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countOfPeaks(tt.args.nums, tt.args.queries); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("countOfPeaks() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
