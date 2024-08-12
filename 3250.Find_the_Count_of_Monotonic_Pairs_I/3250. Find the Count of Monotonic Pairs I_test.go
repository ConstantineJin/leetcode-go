package main

import "testing"

func Test_countOfPairs(t *testing.T) {
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
			args: args{nums: []int{2, 3, 2}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{nums: []int{5, 5, 5, 5}},
			want: 126,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfPairs(tt.args.nums); got != tt.want {
				t.Errorf("countOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
