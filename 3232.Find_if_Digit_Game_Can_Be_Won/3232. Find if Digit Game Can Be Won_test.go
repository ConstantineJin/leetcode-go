package main

import "testing"

func Test_canAliceWin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{nums: []int{1, 2, 3, 4, 10}},
			want: false,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 2, 3, 4, 5, 14}},
			want: true,
		},
		{
			name: "Example3",
			args: args{nums: []int{1, 2, 3, 4, 5, 14}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAliceWin(tt.args.nums); got != tt.want {
				t.Errorf("canAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
