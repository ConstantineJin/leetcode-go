package main

import "testing"

func Test_maxAscendingSum(t *testing.T) {
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
			args:    args{nums: []int{10, 20, 30, 5, 10, 50}},
			wantAns: 65,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{10, 20, 30, 40, 50}},
			wantAns: 150,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{12, 17, 15, 13, 10, 11, 12}},
			wantAns: 33,
		},
		{
			name:    "Example4",
			args:    args{nums: []int{100, 10, 1}},
			wantAns: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxAscendingSum(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maxAscendingSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
