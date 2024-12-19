package main

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
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
			args:    args{arr: []int{4, 3, 2, 1, 0}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{arr: []int{1, 0, 2, 3, 4}},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxChunksToSorted(tt.args.arr); gotAns != tt.wantAns {
				t.Errorf("maxChunksToSorted() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
