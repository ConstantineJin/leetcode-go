package main

import "testing"

func Test_minimumOperations(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 3, 4}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{3, 6, 9}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumOperations(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minimumOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
