package main

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
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
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{11}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfArithmeticSlices(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
