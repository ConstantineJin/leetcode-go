package main

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			wantAns: 5,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 4, 2, 3, 1, 4},
				k:    3,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumLength(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maximumLength() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
