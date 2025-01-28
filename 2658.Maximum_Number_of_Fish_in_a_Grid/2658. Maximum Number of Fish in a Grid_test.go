package main

import "testing"

func Test_findMaxFish(t *testing.T) {
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
			args:    args{grid: [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMaxFish(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("findMaxFish() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
