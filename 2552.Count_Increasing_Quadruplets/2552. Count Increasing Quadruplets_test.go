package main

import "testing"

func Test_countQuadruplets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{1, 3, 2, 4, 5}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3, 4}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countQuadruplets(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countQuadruplets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
