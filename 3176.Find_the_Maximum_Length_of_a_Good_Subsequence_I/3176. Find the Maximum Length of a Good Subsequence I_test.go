package main

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 2, 1, 1, 3},
				k:    2,
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 1},
				k:    0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumLength(tt.args.nums, tt.args.k); gotAns != tt.want {
				t.Errorf("maximumLength() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
