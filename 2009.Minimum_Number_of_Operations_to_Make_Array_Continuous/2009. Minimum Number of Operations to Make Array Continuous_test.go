package main

import "testing"

func Test_minOperations(t *testing.T) {
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
			args: args{nums: []int{4, 2, 5, 3}},
			want: 0,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 2, 3, 5, 6}},
			want: 1,
		},
		{
			name: "Example3",
			args: args{nums: []int{1, 10, 100, 1000}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
