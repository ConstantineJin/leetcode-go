package main

import "testing"

func Test_kSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{2, 4, -2},
				k:    5,
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, -2, 3, 4, -10, 12},
				k:    16,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
