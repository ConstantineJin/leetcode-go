package main

import "testing"

func Test_alternatingSubarray(t *testing.T) {
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
			args:    args{nums: []int{2, 3, 4, 3, 4}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{4, 5, 6}},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{3, 3, 3, 2, 3, 3, 2, 1, 1}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := alternatingSubarray(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("alternatingSubarray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
