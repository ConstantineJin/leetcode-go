package main

import "testing"

func Test_countBadPairs(t *testing.T) {
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
			args: args{nums: []int{4, 1, 3, 3}},
			want: 5,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBadPairs(tt.args.nums); got != tt.want {
				t.Errorf("countBadPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
