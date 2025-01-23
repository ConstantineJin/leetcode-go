package main

import "testing"

func Test_countServers(t *testing.T) {
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
			args:    args{grid: [][]int{{1, 0}, {0, 1}}},
			wantAns: 0,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]int{{1, 0}, {1, 1}}},
			wantAns: 3,
		},
		{
			name:    "Example3",
			args:    args{grid: [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countServers(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("countServers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
