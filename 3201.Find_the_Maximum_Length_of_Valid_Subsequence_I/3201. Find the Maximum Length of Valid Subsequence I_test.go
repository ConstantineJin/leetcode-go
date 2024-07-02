package main

import "testing"

func Test_maximumLength(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 3, 4}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 1, 1, 2, 1, 2}},
			wantAns: 6,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{1, 3}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumLength(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maximumLength() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
