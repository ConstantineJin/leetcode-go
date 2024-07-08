package main

import "testing"

func Test_findMiddleIndex(t *testing.T) {
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
			args: args{nums: []int{2, 3, -1, 8, 4}},
			want: 3,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, -1, 4}},
			want: 2,
		},
		{
			name: "Example3",
			args: args{nums: []int{2, 5}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMiddleIndex(tt.args.nums); got != tt.want {
				t.Errorf("findMiddleIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
