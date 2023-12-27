package main

import "testing"

func Test_isWinner(t *testing.T) {
	type args struct {
		player1 []int
		player2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				player1: []int{4, 10, 7, 9},
				player2: []int{6, 5, 2, 3},
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				player1: []int{3, 5, 7, 6},
				player2: []int{8, 10, 10, 2},
			},
			want: 2,
		},
		{
			name: "Example3",
			args: args{
				player1: []int{2, 3},
				player2: []int{4, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWinner(tt.args.player1, tt.args.player2); got != tt.want {
				t.Errorf("isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
