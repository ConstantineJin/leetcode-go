package main

import "testing"

func Test_semiOrderedPermutation(t *testing.T) {
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
			args: args{nums: []int{2, 1, 4, 3}},
			want: 2,
		},
		{
			name: "Example2",
			args: args{nums: []int{2, 4, 1, 3}},
			want: 3,
		},
		{
			name: "Example3",
			args: args{nums: []int{1, 3, 4, 2, 5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := semiOrderedPermutation(tt.args.nums); got != tt.want {
				t.Errorf("semiOrderedPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
