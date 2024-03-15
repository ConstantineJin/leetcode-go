package main

import "testing"

func Test_sellingWood(t *testing.T) {
	type args struct {
		m      int
		n      int
		prices [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				m:      3,
				n:      5,
				prices: [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}},
			},
			want: 19,
		},
		{
			name: "Example2",
			args: args{
				m:      4,
				n:      6,
				prices: [][]int{{3, 2, 10}, {1, 4, 2}, {4, 1, 3}},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sellingWood(tt.args.m, tt.args.n, tt.args.prices); got != tt.want {
				t.Errorf("sellingWood() = %v, want %v", got, tt.want)
			}
		})
	}
}
