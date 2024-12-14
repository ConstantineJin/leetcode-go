package main

import "testing"

func Test_continuousSubarrays(t *testing.T) {
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
			args:    args{nums: []int{5, 4, 2, 4}},
			wantAns: 8,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3}},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := continuousSubarrays(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("continuousSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
