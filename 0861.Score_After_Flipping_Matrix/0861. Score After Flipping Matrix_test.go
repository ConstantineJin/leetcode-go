package main

import "testing"

func Test_matrixScore(t *testing.T) {
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
			args:    args{grid: [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}},
			wantAns: 39,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]int{{0}}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := matrixScore(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("matrixScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
