package main

import "testing"

func Test_minimumSize(t *testing.T) {
	type args struct {
		nums          []int
		maxOperations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums:          []int{9},
				maxOperations: 2,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				nums:          []int{2, 4, 8, 2},
				maxOperations: 4,
			},
			want: 2,
		},
		{
			name: "Example3",
			args: args{
				nums:          []int{7, 17},
				maxOperations: 2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSize(tt.args.nums, tt.args.maxOperations); got != tt.want {
				t.Errorf("minimumSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
