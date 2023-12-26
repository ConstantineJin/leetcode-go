package main

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
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
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			wantAns: [][]int{{7}, {2, 2, 3}},
		},
		{
			name: "Example2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			wantAns: [][]int{{3, 5}, {2, 3, 3}, {2, 2, 2, 2}},
		},
		{
			name: "Example3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combinationSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
