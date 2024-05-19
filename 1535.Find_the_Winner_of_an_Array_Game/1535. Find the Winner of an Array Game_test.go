package main

import "testing"

func Test_getWinner(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				arr: []int{2, 1, 3, 5, 4, 6, 7},
				k:   2,
			},
			want: 5,
		},
		{
			name: "Example2",
			args: args{
				arr: []int{3, 2, 1},
				k:   10,
			},
			want: 3,
		},
		{
			name: "Example3",
			args: args{
				arr: []int{1, 9, 8, 2, 3, 7, 6, 4, 5},
				k:   7,
			},
			want: 9,
		},
		{
			name: "Example4",
			args: args{
				arr: []int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99},
				k:   1000000000,
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWinner(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("getWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
