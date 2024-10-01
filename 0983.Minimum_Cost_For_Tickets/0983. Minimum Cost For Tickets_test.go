package main

import "testing"

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				days:  []int{1, 4, 6, 7, 8, 20},
				costs: []int{2, 7, 15},
			},
			want: 11,
		},
		{
			name: "Example2",
			args: args{
				days:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
				costs: []int{2, 7, 15},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}
