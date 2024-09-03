package main

import "testing"

func Test_maxStrength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{nums: []int{3, -1, -5, 2, 5, -9}},
			want: 1350,
		},
		{
			name: "Example2",
			args: args{nums: []int{-4, -5, -4}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxStrength(tt.args.nums); got != tt.want {
				t.Errorf("maxStrength() = %v, want %v", got, tt.want)
			}
		})
	}
}
