package main

import "testing"

func Test_findGCD(t *testing.T) {
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
			args: args{nums: []int{2, 5, 6, 9, 10}},
			want: 2,
		},
		{
			name: "Example2",
			args: args{nums: []int{7, 5, 6, 8, 3}},
			want: 1,
		},
		{
			name: "Example3",
			args: args{nums: []int{3, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGCD(tt.args.nums); got != tt.want {
				t.Errorf("findGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
