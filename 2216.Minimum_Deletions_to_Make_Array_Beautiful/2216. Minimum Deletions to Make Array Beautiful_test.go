package main

import "testing"

func Test_minDeletion(t *testing.T) {
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
			args:    args{nums: []int{1, 1, 2, 3, 5}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 1, 2, 2, 3, 3}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minDeletion(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minDeletion() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
