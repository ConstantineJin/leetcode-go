package main

import "testing"

func Test_rootCount(t *testing.T) {
	type args struct {
		edges   [][]int
		guesses [][]int
		k       int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				edges:   [][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}},
				guesses: [][]int{{1, 3}, {0, 1}, {1, 0}, {2, 4}},
				k:       3,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				edges:   [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
				guesses: [][]int{{1, 0}, {3, 4}, {2, 1}, {3, 2}},
				k:       1,
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := rootCount(tt.args.edges, tt.args.guesses, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("rootCount() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
