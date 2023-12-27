package main

import "testing"

func Test_findDuplicate(t *testing.T) {
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
			args:    args{nums: []int{1, 3, 4, 2, 2}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{3, 1, 3, 4, 2}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findDuplicate(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findDuplicate() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
