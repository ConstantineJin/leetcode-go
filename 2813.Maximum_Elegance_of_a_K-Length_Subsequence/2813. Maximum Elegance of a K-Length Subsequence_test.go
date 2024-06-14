package main

import "testing"

func Test_findMaximumElegance(t *testing.T) {
	type args struct {
		items [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				items: [][]int{{3, 2}, {5, 1}, {10, 1}},
				k:     2,
			},
			want: 17,
		},
		{
			name: "Example2",
			args: args{
				items: [][]int{{3, 1}, {3, 1}, {2, 2}, {5, 3}},
				k:     3,
			},
			want: 19,
		},
		{
			name: "Example3",
			args: args{
				items: [][]int{{1, 1}, {2, 1}, {3, 1}},
				k:     3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumElegance(tt.args.items, tt.args.k); got != tt.want {
				t.Errorf("findMaximumElegance() = %v, want %v", got, tt.want)
			}
		})
	}
}
