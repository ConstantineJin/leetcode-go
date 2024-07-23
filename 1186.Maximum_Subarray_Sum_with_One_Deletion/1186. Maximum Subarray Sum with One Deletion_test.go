package main

import "testing"

func Test_maximumSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{arr: []int{1, -2, 3, 0}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{arr: []int{1, -2, -2, 3}},
			want: 3,
		},
		{
			name: "Example3",
			args: args{arr: []int{-1, -1, -1, -1}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSum(tt.args.arr); got != tt.want {
				t.Errorf("maximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
