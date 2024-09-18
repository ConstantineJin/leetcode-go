package main

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{10, 2}},
			wantAns: "210",
		},
		{
			name:    "Example2",
			args:    args{nums: []int{3, 30, 34, 5, 9}},
			wantAns: "9534330",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := largestNumber(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("largestNumber() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
