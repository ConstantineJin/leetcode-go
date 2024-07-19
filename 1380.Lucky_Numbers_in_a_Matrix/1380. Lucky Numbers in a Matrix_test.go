package main

import (
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{matrix: [][]int{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := luckyNumbers(tt.args.matrix); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("luckyNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
