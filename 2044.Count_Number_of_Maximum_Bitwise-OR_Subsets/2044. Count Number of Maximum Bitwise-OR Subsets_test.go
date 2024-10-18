package main

import "testing"

func Test_countMaxOrSubsets(t *testing.T) {
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
			args:    args{nums: []int{3, 1}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 2, 2}},
			wantAns: 7,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{3, 2, 1, 5}},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countMaxOrSubsets(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countMaxOrSubsets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
