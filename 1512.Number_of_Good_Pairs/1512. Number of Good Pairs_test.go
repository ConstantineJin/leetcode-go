package main

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 3, 1, 1, 3}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 1, 1, 1}},
			wantAns: 6,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{1, 2, 3}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numIdenticalPairs(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("numIdenticalPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
