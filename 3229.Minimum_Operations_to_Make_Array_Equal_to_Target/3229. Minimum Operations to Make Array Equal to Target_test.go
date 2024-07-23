package main

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums   []int
		target []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				nums:   []int{3, 5, 1, 2},
				target: []int{4, 6, 2, 4},
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				nums:   []int{1, 3, 2},
				target: []int{2, 1, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
