package main

import "testing"

func Test_maxFrequencyElements(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 2, 3, 1, 4}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3, 4, 5}},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxFrequencyElements(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maxFrequencyElements() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
