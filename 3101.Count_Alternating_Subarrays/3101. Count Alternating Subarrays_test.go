package main

import "testing"

func Test_countAlternatingSubarrays(t *testing.T) {
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
			args:    args{nums: []int{0, 1, 1, 1}},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 0, 1, 0}},
			wantAns: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countAlternatingSubarrays(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countAlternatingSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
