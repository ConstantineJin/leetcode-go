package main

import "testing"

func Test_findValueOfPartition(t *testing.T) {
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
			args: args{nums: []int{1, 3, 2, 4}},
			want: 1,
		},
		{
			name: "Example2",
			args: args{nums: []int{100, 1, 10}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findValueOfPartition(tt.args.nums); got != tt.want {
				t.Errorf("findValueOfPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
