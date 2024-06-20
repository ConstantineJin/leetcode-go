package main

import "testing"

func Test_countBeautifulPairs(t *testing.T) {
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
			args:    args{nums: []int{2, 5, 1, 4}},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{11, 21, 12}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countBeautifulPairs(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countBeautifulPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
