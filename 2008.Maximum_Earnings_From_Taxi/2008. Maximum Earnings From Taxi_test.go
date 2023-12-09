package main

import "testing"

func Test_maxTaxiEarnings(t *testing.T) {
	type args struct {
		n     int
		rides [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				n:     5,
				rides: [][]int{{2, 5, 4}, {1, 5, 1}},
			},
			want: 7,
		},
		{
			name: "Example2",
			args: args{
				n:     20,
				rides: [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTaxiEarnings(tt.args.n, tt.args.rides); got != tt.want {
				t.Errorf("maxTaxiEarnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
