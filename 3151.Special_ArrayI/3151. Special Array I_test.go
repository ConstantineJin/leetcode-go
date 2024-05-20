package main

import "testing"

func Test_isArraySpecial(t *testing.T) {
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
			args: args{nums: []int{1}},
			want: true,
		},
		{
			name: "Example2",
			args: args{nums: []int{2, 1, 4}},
			want: true,
		},
		{
			name: "Example3",
			args: args{nums: []int{4, 3, 1, 6}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums); got != tt.want {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
