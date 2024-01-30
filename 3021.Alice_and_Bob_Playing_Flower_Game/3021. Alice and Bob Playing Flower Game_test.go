package main

import "testing"

func Test_flowerGame(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				n: 3,
				m: 2,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				n: 1,
				m: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flowerGame(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("flowerGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
