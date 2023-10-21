package main

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		{
			name: "Example1",
			args: args{
				n:     3,
				edges: [][]int{{0, 1}, {0, 2}, {1, 2}},
			},
			wantRes: 0,
		},
		{
			name: "Example2",
			args: args{
				n:     7,
				edges: [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}},
			},
			wantRes: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countPairs(tt.args.n, tt.args.edges); gotRes != tt.wantRes {
				t.Errorf("countPairs() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
