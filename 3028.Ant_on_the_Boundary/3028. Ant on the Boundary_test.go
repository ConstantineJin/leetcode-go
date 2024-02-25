package main

import "testing"

func Test_returnToBoundaryCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{2, 3, -5}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{3, 2, -3, -4}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := returnToBoundaryCount(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("returnToBoundaryCount() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
