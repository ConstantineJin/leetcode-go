package main

import "testing"

func Test_canTraverseAllPairs(t *testing.T) {
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
			args: args{nums: []int{2, 3, 6}},
			want: true,
		},
		{
			name: "Example2",
			args: args{nums: []int{3, 9, 5}},
			want: false,
		},
		{
			name: "Example3",
			args: args{nums: []int{4, 3, 12, 8}},
			want: true,
		},
		{
			name: "Example4",
			args: args{nums: []int{28, 39}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canTraverseAllPairs(tt.args.nums); got != tt.want {
				t.Errorf("canTraverseAllPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
