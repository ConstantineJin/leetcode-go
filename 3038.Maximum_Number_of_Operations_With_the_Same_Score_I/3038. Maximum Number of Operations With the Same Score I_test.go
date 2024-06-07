package main

import "testing"

func Test_maxOperations(t *testing.T) {
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
			args:    args{nums: []int{3, 2, 1, 4, 5}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{3, 2, 6, 1, 4}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxOperations(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maxOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
