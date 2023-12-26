package main

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Example1",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			wantAns: [][]int{{2, 6}, {1, 7}, {1, 2, 5}, {1, 1, 6}},
		},
		{
			name: "Example2",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			wantAns: [][]int{{5}, {1, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combinationSum2() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
