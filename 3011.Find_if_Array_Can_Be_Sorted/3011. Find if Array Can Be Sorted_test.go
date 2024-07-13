package main

import "testing"

func Test_canSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{nums: []int{8, 4, 2, 30, 15}},
			want: true,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "Example3",
			args: args{nums: []int{3, 16, 8, 4, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSortArray(tt.args.nums); got != tt.want {
				t.Errorf("canSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
