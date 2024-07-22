package main

import "testing"

func Test_losingPlayer(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				x: 2,
				y: 7,
			},
			want: "Alice",
		},
		{
			name: "Example2",
			args: args{
				x: 4,
				y: 11,
			},
			want: "Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := losingPlayer(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("losingPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
