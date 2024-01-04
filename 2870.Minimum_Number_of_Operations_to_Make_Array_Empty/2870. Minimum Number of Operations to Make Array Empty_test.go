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
			args:    args{nums: []int{2, 3, 3, 2, 2, 4, 2, 3, 4}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 1, 2, 2, 3, 3}},
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
