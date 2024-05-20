package main

import "testing"

func Test_sumDigitDifferences(t *testing.T) {
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
			args: args{nums: []int{13, 23, 12}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{nums: []int{10, 10, 10, 10}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitDifferences(tt.args.nums); got != tt.want {
				t.Errorf("sumDigitDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
