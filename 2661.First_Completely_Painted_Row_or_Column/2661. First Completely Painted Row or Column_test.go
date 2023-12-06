package main

import "testing"

func Test_firstCompleteIndex(t *testing.T) {
	type args struct {
		arr []int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				arr: []int{1, 3, 4, 2},
				mat: [][]int{{1, 4}, {2, 3}},
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				arr: []int{2, 8, 7, 4, 1, 3, 5, 6, 9},
				mat: [][]int{{3, 2, 4}, {1, 4, 6}, {8, 7, 9}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstCompleteIndex(tt.args.arr, tt.args.mat); got != tt.want {
				t.Errorf("firstCompleteIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
