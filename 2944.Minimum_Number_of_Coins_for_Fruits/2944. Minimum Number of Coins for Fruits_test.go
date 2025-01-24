package main

import "testing"

func Test_minimumCoins(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{prices: []int{3, 1, 2}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{prices: []int{1, 10, 1, 1}},
			want: 2,
		},
		{
			name: "Example3",
			args: args{prices: []int{26, 18, 6, 12, 49, 7, 45, 45}},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCoins(tt.args.prices); got != tt.want {
				t.Errorf("minimumCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
