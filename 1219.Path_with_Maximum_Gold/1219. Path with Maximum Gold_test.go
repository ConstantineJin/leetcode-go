package main

import "testing"

func Test_getMaximumGold(t *testing.T) {
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
			args:    args{grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}},
			wantAns: 24,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}},
			wantAns: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getMaximumGold(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("getMaximumGold() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
