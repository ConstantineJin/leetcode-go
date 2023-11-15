package main

import "testing"

func Test_maximizeSum(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3, 4, 5}, k: 3},
			want: 18,
		},
		{
			name: "Example2",
			args: args{nums: []int{5, 5, 5}, k: 2},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximizeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
