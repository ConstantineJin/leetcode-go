package main

import "testing"

func Test_orangesRotting(t *testing.T) {
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
			args:    args{grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}},
			wantAns: -1,
		},
		{
			name:    "Example3",
			args:    args{grid: [][]int{{0, 2}}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := orangesRotting(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("orangesRotting() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
