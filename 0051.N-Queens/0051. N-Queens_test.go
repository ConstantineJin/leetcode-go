package main

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]string
	}{
		{
			name:    "Example1",
			args:    args{n: 4},
			wantAns: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			name:    "Example2",
			args:    args{n: 1},
			wantAns: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := solveNQueens(tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("solveNQueens() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
