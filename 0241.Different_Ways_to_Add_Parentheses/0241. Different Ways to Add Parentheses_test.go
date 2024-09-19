package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name:    "Example1",
			args:    args{expression: "2-1-1"},
			wantAns: []int{0, 2},
		},
		{
			name:    "Example2",
			args:    args{expression: "2*3-4*5"},
			wantAns: []int{-34, -14, -10, -10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := diffWaysToCompute(tt.args.expression)
			sort.Ints(gotAns)
			sort.Ints(tt.wantAns)
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("diffWaysToCompute() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
