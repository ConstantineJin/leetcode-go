package main

import "testing"

func Test_checkMove(t *testing.T) {
	type args struct {
		board [][]byte
		rMove int
		cMove int
		color byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				board: [][]byte{{'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}},
				rMove: 4,
				cMove: 3,
				color: 'B',
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				board: [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', 'B', '.', '.', 'W', '.', '.', '.'}, {'.', '.', '.', 'W', 'B', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', 'B', 'W', '.', '.'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', 'W', '.'}, {'.', '.', '.', '.', '.', '.', '.', 'B'}},
				rMove: 4,
				cMove: 4,
				color: 'W',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkMove(tt.args.board, tt.args.rMove, tt.args.cMove, tt.args.color); got != tt.want {
				t.Errorf("checkMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
