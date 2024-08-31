package main

import "testing"

func Test_canMakeSquare(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{grid: [][]byte{{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'B'}}},
			want: true,
		},
		{
			name: "Example2",
			args: args{grid: [][]byte{{'B', 'W', 'B'}, {'W', 'B', 'W'}, {'B', 'W', 'B'}}},
			want: false,
		},
		{
			name: "Example3",
			args: args{grid: [][]byte{{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'W'}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeSquare(tt.args.grid); got != tt.want {
				t.Errorf("canMakeSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
