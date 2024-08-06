package main

import "testing"

func Test_winningPlayerCount(t *testing.T) {
	type args struct {
		n    int
		pick [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				n:    4,
				pick: [][]int{{0, 0}, {1, 0}, {1, 0}, {2, 1}, {2, 1}, {2, 0}},
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				n:    5,
				pick: [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}},
			},
			wantAns: 0,
		},
		{
			name: "Example3",
			args: args{
				n:    4,
				pick: [][]int{{1, 1}, {2, 4}, {2, 4}, {2, 4}},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := winningPlayerCount(tt.args.n, tt.args.pick); gotAns != tt.wantAns {
				t.Errorf("winningPlayerCount() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
