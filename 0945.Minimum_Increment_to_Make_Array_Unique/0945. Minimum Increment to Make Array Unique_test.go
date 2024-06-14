package main

import "testing"

func Test_minIncrementForUnique(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 2}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{3, 2, 1, 2, 1, 7}},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minIncrementForUnique(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minIncrementForUnique() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
