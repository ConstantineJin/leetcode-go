package main

import "testing"

func Test_minOperations(t *testing.T) {
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
			args:    args{nums: []int{0, 1, 1, 1, 0, 0}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{0, 1, 1, 1}},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minOperations(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
