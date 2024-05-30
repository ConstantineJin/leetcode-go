package main

import "testing"

func Test_countTriplets(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{arr: []int{2, 3, 1, 6, 7}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{arr: []int{1, 1, 1, 1, 1}},
			wantAns: 10,
		},
		{
			name:    "Example3",
			args:    args{arr: []int{2, 3}},
			wantAns: 0,
		},
		{
			name:    "Example4",
			args:    args{arr: []int{1, 3, 5, 7, 9}},
			wantAns: 3,
		},
		{
			name:    "Example5",
			args:    args{arr: []int{7, 11, 12, 9, 5, 2, 7, 17, 22}},
			wantAns: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countTriplets(tt.args.arr); gotAns != tt.wantAns {
				t.Errorf("countTriplets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
