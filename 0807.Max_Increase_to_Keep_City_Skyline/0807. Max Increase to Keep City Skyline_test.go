package main

import "testing"

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
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
			args:    args{grid: [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}},
			wantAns: 35,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxIncreaseKeepingSkyline(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("maxIncreaseKeepingSkyline() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
