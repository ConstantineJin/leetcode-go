package main

import "testing"

func Test_duplicateNumbersXOR(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 1, 3}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3}},
			wantAns: 0,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{1, 2, 2, 1}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := duplicateNumbersXOR(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("duplicateNumbersXOR() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
