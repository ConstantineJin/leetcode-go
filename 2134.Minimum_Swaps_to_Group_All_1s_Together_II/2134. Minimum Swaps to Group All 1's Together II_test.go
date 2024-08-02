package main

import "testing"

func Test_minSwaps(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{nums: []int{0, 1, 0, 1, 1, 0, 0}},
			want: 1,
		},
		{
			name: "Example2",
			args: args{nums: []int{0, 1, 1, 1, 0, 0, 1, 1, 0}},
			want: 2,
		},
		{
			name: "Example3",
			args: args{nums: []int{1, 1, 0, 0, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.nums); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
