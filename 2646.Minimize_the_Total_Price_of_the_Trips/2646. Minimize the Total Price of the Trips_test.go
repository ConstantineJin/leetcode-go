package main

import "testing"

func Test_minimumTotalPrice(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {1, 2}, {1, 3}},
				price: []int{2, 2, 10, 6},
				trips: [][]int{{0, 3}, {2, 1}, {2, 3}},
			},
			want: 23,
		},
		{
			name: "Example2",
			args: args{
				n:     2,
				edges: [][]int{{0, 1}},
				price: []int{2, 2},
				trips: [][]int{{0, 0}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotalPrice(tt.args.n, tt.args.edges, tt.args.price, tt.args.trips); got != tt.want {
				t.Errorf("minimumTotalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
