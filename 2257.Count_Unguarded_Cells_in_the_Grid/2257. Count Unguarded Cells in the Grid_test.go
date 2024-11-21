package main

import "testing"

func Test_countUnguarded(t *testing.T) {
	type args struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				m:      4,
				n:      6,
				guards: [][]int{{0, 0}, {1, 1}, {2, 3}},
				walls:  [][]int{{0, 1}, {2, 2}, {1, 4}},
			},
			wantAns: 7,
		},
		{
			name: "Example2",
			args: args{
				m:      3,
				n:      3,
				guards: [][]int{{1, 1}},
				walls:  [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}},
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countUnguarded(tt.args.m, tt.args.n, tt.args.guards, tt.args.walls); gotAns != tt.wantAns {
				t.Errorf("countUnguarded() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
