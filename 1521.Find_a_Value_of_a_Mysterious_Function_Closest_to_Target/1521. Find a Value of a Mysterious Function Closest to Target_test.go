package main

import "testing"

func Test_closestToTarget(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				arr:    []int{9, 12, 3, 7, 15},
				target: 5,
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				arr:    []int{1000000, 1000000, 1000000},
				target: 1,
			},
			want: 999999,
		},
		{
			name: "Example1",
			args: args{
				arr:    []int{1, 2, 4, 8, 16},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestToTarget(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("closestToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
