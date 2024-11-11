package main

import "testing"

func Test_primeSubOperation(t *testing.T) {
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
			args: args{nums: []int{4, 9, 6, 10}},
			want: true,
		},
		{
			name: "Example2",
			args: args{nums: []int{6, 8, 11, 12}},
			want: true,
		},
		{
			name: "Example3",
			args: args{nums: []int{5, 8, 3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primeSubOperation(tt.args.nums); got != tt.want {
				t.Errorf("primeSubOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
