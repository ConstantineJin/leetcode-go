package main

import "testing"

func Test_stoneGameII(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{piles: []int{2, 7, 9, 4, 4}},
			want: 10,
		},
		{
			name: "Example2",
			args: args{piles: []int{1, 2, 3, 4, 5, 100}},
			want: 104,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameII(tt.args.piles); got != tt.want {
				t.Errorf("stoneGameII() = %v, want %v", got, tt.want)
			}
		})
	}
}
