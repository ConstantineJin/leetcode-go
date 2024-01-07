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
			args:    args{nums: []int{2, 4, 6, 8, 10}},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{7, 7, 7, 7, 7}},
			wantAns: 16,
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
