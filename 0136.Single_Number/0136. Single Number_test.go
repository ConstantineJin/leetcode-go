package main

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantSingle int
	}{
		{
			name:       "Example1",
			args:       args{nums: []int{2, 2, 1}},
			wantSingle: 1,
		},
		{
			name:       "Example2",
			args:       args{nums: []int{4, 1, 2, 1, 2}},
			wantSingle: 4,
		},
		{
			name:       "Example3",
			args:       args{nums: []int{1}},
			wantSingle: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSingle := singleNumber(tt.args.nums); gotSingle != tt.wantSingle {
				t.Errorf("singleNumber() = %v, want %v", gotSingle, tt.wantSingle)
			}
		})
	}
}
