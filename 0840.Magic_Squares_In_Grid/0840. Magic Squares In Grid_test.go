package main

import "testing"

func Test_numMagicSquaresInside(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{grid: [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]int{{8}}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numMagicSquaresInside(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("numMagicSquaresInside() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
