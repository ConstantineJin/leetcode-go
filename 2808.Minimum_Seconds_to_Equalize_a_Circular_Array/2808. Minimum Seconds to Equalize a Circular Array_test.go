package main

import "testing"

func Test_minimumSeconds(t *testing.T) {
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
			args: args{nums: []int{1, 2, 1, 2}},
			want: 1,
		},
		{
			name: "Example2",
			args: args{nums: []int{2, 1, 3, 3, 2}},
			want: 2,
		},
		{
			name: "Example3",
			args: args{nums: []int{5, 5, 5, 5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSeconds(tt.args.nums); got != tt.want {
				t.Errorf("minimumSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
