package main

import "testing"

func Test_maximumSumOfHeights(t *testing.T) {
	type args struct {
		maxHeights []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{maxHeights: []int{5, 3, 4, 1, 1}},
			want: 13,
		},
		{
			name: "Example2",
			args: args{maxHeights: []int{6, 5, 3, 9, 2, 7}},
			want: 22,
		},
		{
			name: "Example3",
			args: args{maxHeights: []int{3, 2, 5, 5, 2, 3}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSumOfHeights(tt.args.maxHeights); got != tt.want {
				t.Errorf("maximumSumOfHeights() = %v, want %v", got, tt.want)
			}
		})
	}
}
